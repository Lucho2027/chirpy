export interface CreateChirpArgs {
    message: string;
    userId: string;
}
export interface CreateChirpRow {
    id: string;
    message: string;
    userId: string;
    createdAt: Date | null;
    updatedAt: Date | null;
}
export interface GetAllChirpsRow {
    id: string;
    message: string;
    userId: string;
    createdAt: Date | null;
    updatedAt: Date | null;
}
export interface GetChirpByIdArgs {
    id: string;
}
export interface GetChirpByIdRow {
    id: string;
    message: string;
    userId: string;
    createdAt: Date | null;
    updatedAt: Date | null;
}
export interface DeleteChirpByIdArgs {
    chirpId: string;
    userId: string;
}
export interface GetAllChirpsByAuthorArgs {
    userId: string;
}
export interface GetAllChirpsByAuthorRow {
    id: string;
    message: string;
    userId: string;
    createdAt: Date | null;
    updatedAt: Date | null;
}
