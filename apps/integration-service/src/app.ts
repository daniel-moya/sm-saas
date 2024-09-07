import * as grpc from '@grpc/grpc-js';
import {
    IntegrationServiceService,
} from './proto/integration';
import { IntegrationServiceImpl } from './services/integration-service';

class App {
    start() {
        const server = new grpc.Server();
        server.addService(IntegrationServiceService, IntegrationServiceImpl);

        server.bindAsync(
            '127.0.0.1:50051',
            grpc.ServerCredentials.createInsecure(),
            () => {
            }
        );
    }
}

export default App

