export interface Room {
    id: string;
    room_code: string;
    owner_id: string;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
}

export interface RoomForm {
    name: string;
    description: string;
}

export interface FetchRoomResponse {
    data: Room;
    message: string;
    status: string;
}

export interface FetchRoomListResponse {
    data: {
        pagination: {
            limit: number;
            page: number;
            page_total: number;
        };
        rooms: Room[];
    };
    message: string;
    status: string;
}

export interface CreateRoomResponse {
    data: Room;
    message: string;
    status: string;
}

export interface FetchJoinResponse {
    message: string,
    status: string
}

export interface LeaveRoomResponse {
    message: string;
    status: string;
}

export interface DeleteRoomResponse {
    message: string;
    status: string;
}

export interface Member {
    room_id: string,
    user_id: string,
    username: string,
    profile_url: string,
    joined_at: string
    online?: boolean;
}

export interface GetMembersResponse {
    data: Member[];
    message: string;
    status: string;
}