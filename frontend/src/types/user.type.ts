export interface User {
  id: number;
  username: string;
  profile_url: string;
  created_at: string;
}

export interface FetchMeResponse {
    data: User;
    message: string;
    status: string;
}