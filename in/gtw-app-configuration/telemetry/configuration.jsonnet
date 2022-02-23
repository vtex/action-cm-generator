{
  openTelemetry: {
    enableTracing: true,
    enableDebugMode: true,
    sampling: {
      maxConcurrentTraces: 100,
      perAccountConfig: {
        default: {
          defaultProbability: 0.01,
          perOperationProbability: {
            'GET healthcheck': 0,
            'GET metrics': 0,
            'GET api/pvt/installments': 0.003,
            'GET api/pvt/merchants/payment-systems': 0.01,
            'GET api/pvt/admin/installments/provider': 0.003,
            'POST api/pvt/merchantredirectitem': 0.003,
            'POST api/pvt/transactions/{transactionId}/authorization-request': 0.075,
            'POST api/pvt/transactions/{transactionId}/settlement-request': 0.075,
            'POST api/pvt/transactions/{transactionId}/cancellation-request': 0.075,
            'POST api/pvt/transactions/{transactionId}/refunding-request': 0.075,
            'HandleEvent PostBackStatusWorker': 1,
          },
        },
      },
    },
    clientHeadersTracking: {
      trackedRequestHeaders: [],
      trackedResponseHeaders: [],
    },
    serverHeadersTracking: {
      trackedRequestHeaders: ['x-vtex-user-agent'],
      trackedResponseHeaders: [],
    },
  },
}
