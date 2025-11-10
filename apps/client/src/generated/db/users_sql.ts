export interface CreateUserArgs {
    email: string;
    password: string;
}
export interface CreateUserRow {
    id: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    email: string;
    hashedPassword: string;
    isChirpyRed: boolean;
}
export interface GetByEmailArgs {
    email: string;
}
export interface GetByEmailRow {
    email: string;
    hashedPassword: string;
    id: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    isChirpyRed: boolean;
}
export interface UpdateUserArgs {
    email: string;
    password: string;
    id: string;
}
export interface UpdateUserRow {
    id: string;
    createdAt: Date | null;
    updatedAt: Date | null;
    email: string;
    hashedPassword: string;
    isChirpyRed: boolean;
}
export interface UpgradeUserArgs {
    isChirpyRed: boolean;
    id: string;
}
export interface UpgradeUserRow {
    id: string;
}
