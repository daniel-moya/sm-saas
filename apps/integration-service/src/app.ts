import * as grpc from '@grpc/grpc-js';
import { IntegrationService } from './proto/integration';

class App {
    start() {
        // Add the implemented service to the server
        //server.addService(IntegrationServiceProto.IntegrationService.service, {
        //    StoreUserTokens: storeUserTokens,
        //});

        const PORT = process.env.PORT || '50051';

        const server = new grpc.Server();

        server.bindAsync(
            '127.0.0.1:50051',
            grpc.ServerCredentials.createInsecure(),
            () => {
                server.start();
            }
        );
    }
}

export default App

