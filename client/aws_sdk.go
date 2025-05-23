package client

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/rs/zerolog"

	"github.com/OpsHelmInc/cloudquery/v2/client/spec"
)

func ConfigureAwsSDK(ctx context.Context, logger zerolog.Logger, awsPluginSpec *spec.Spec, account spec.Account, stsClient AssumeRoleAPIClient) (aws.Config, error) {
	var err error
	var awsCfg aws.Config

	// This sets MaxRetries & MaxBackoff, too
	awsPluginSpec.SetDefaults()

	configFns := []func(*config.LoadOptions) error{
		config.WithDefaultRegion(defaultRegion),
		// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/retries-timeouts/
		config.WithRetryer(func() aws.Retryer {
			return retry.NewStandard(func(so *retry.StandardOptions) {
				so.MaxAttempts = *awsPluginSpec.MaxRetries
				so.MaxBackoff = time.Duration(*awsPluginSpec.MaxBackoff) * time.Second
				so.RateLimiter = &NoRateLimiter{}
			})
		}),
	}

	if account.DefaultRegion != "" {
		// According to the docs: If multiple WithDefaultRegion calls are made, the last call overrides the previous call values
		configFns = append(configFns, config.WithDefaultRegion(account.DefaultRegion))
	}

	if account.LocalProfile != "" {
		configFns = append(configFns, config.WithSharedConfigProfile(account.LocalProfile))
	}

	awsCfg, err = config.LoadDefaultConfig(ctx, configFns...)

	if err != nil {
		logger.Error().Err(err).Msg("error loading default config")
		return awsCfg, err
	}

	if awsPluginSpec.EndpointURL != "" {
		awsCfg.BaseEndpoint = &awsPluginSpec.EndpointURL
	}

	if account.RoleARN != "" {
		opts := make([]func(*stscreds.AssumeRoleOptions), 0, 1)
		if account.ExternalID != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.ExternalID = &account.ExternalID
			})
		}
		if account.RoleSessionName != "" {
			opts = append(opts, func(opts *stscreds.AssumeRoleOptions) {
				opts.RoleSessionName = account.RoleSessionName
			})
		}

		if stsClient == nil {
			stsClient = sts.NewFromConfig(awsCfg)
		}
		provider := stscreds.NewAssumeRoleProvider(stsClient, account.RoleARN, opts...)

		awsCfg.Credentials = aws.NewCredentialsCache(provider, func(options *aws.CredentialsCacheOptions) {
			// ExpiryWindow will allow the credentials to trigger refreshing prior to
			// the credentials actually expiring. This is beneficial so race conditions
			// with expiring credentials do not cause requests to fail unexpectedly
			// due to ExpiredToken exceptions.
			//
			// An ExpiryWindow of 5 minute would cause calls to IsExpired() to return true
			// 5 minutes before the credentials are actually expired. This can cause an
			// increased number of requests to refresh the credentials to occur. We balance this with jitter.
			options.ExpiryWindow = 5 * time.Minute
			// Jitter is added to avoid the thundering herd problem of many refresh requests
			// happening all at once.
			options.ExpiryWindowJitterFrac = 0.5
		})
	}

	if awsPluginSpec.AWSDebug {
		awsCfg.ClientLogMode = aws.LogRequestWithBody | aws.LogResponseWithBody | aws.LogRetries
		awsCfg.Logger = AwsLogger{logger.With().Str("accountName", account.AccountName).Logger()}
	}

	// Test out retrieving credentials
	if _, err := awsCfg.Credentials.Retrieve(ctx); err != nil {
		logger.Error().Err(err).Msg("error retrieving credentials")
		return awsCfg, errRetrievingCredentials
	}

	return awsCfg, err
}
