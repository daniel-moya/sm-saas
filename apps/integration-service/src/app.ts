import * as grpc from '@grpc/grpc-js';
import {
    IntegrationServiceService,
    LinkSocialAccountRequest,
    LinkSocialAccountResponse,
    UnlinkSocialAccountRequest,
    UnlinkSocialAccountResponse
} from './proto/integration';

// Implement the actual service methods
const integrationServiceImpl = {
    linkSocialAccount: (
        call: grpc.ServerUnaryCall<LinkSocialAccountRequest, LinkSocialAccountResponse>,
        callback: grpc.sendUnaryData<LinkSocialAccountResponse>,
    ) => {
        const response = LinkSocialAccountResponse.create({ message: "Linked" });
        callback(null, response);
    },

    unlinkSocialAccount: (
        call: grpc.ServerUnaryCall<UnlinkSocialAccountRequest, UnlinkSocialAccountResponse>,
        callback: grpc.sendUnaryData<UnlinkSocialAccountResponse>,
    ) => {
        const response = UnlinkSocialAccountResponse.create({ message: "Unlinked" });
        callback(null, response);
    }
};


class App {
    start() {
        const server = new grpc.Server();
        server.addService(IntegrationServiceService, integrationServiceImpl);

        server.bindAsync(
            '127.0.0.1:50051',
            grpc.ServerCredentials.createInsecure(),
            () => {
                //server.start();
            }
        );
    }
}

export default App

