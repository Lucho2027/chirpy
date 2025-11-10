export interface CreateRefreshTokenArgs {
    token: string;
    userId: string;
    expiresAt: Date | null;
}
export interface CreateRefreshTokenRow {
    token: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    userId: string;
    expiresAt: Date | null;
    revokedAt: Date | null;
}
export interface GetUserFromRefreshTokenArgs {
    token: string;
}
export interface GetUserFromRefreshTokenRow {
    userId: string;
}
export interface RevokeTokenArgs {
    token: string;
}
