/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: 'server',
      removal: input?.stage === 'production' ? 'retain' : 'remove',
      home: 'aws',
    };
  },
  async run() {
    const api = new sst.aws.ApiGatewayV2('url-shortner');
    api.route('POST /', {
      handler: 'bootstrap',
      runtime: 'provided.al2023',
      architecture: 'x86_64',
      bundle: './packages/functions/create-url/build',
      memory: '128 MB',
      logging: {
        retention: '3 days',
      },
    });

    console.log(api.url);
  },
});
