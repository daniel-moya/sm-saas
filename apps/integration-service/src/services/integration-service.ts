
import * as grpc from '@grpc/grpc-js';
import {
    LinkSocialAccountRequest,
    LinkSocialAccountResponse,
    UnlinkSocialAccountRequest,
    UnlinkSocialAccountResponse
} from '../proto/integration';

// Implement the actual service methods
export const IntegrationServiceImpl = {
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
