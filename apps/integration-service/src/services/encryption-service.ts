import bcrypt from "bcrypt";

class EncryptionService {
    async encryptToken(token: string): Promise<string> {
        try {
            const salt = await bcrypt.genSalt(10);
            const hashedToken = await bcrypt.hash(token, salt);
            return hashedToken;
        } catch (err: any) {
            throw Error("Error encrypting password", err)
        }
    }

    async compareToken(token: string, hashedToken: string): Promise<boolean> {
        try {
            return await bcrypt.compare(token, hashedToken);
        } catch (err: any) {
            throw Error("Error comparing/validating password hashes", err)
        }
    }
}

export default EncryptionService
