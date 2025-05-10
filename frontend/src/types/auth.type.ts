export interface RegisterForm {
  username: string;
  password: string;
  confirm_password?: string;
}

export interface LoginForm {
  username: string;
  password: string;
}

export interface FetchLoginResponse {
    data: {
      username: string;
      profile_url: string;
      created_at: string;
      access_token: string;
      refresh_token: string;
    };
    message: string;
    status: string;
}

export interface FetchRegisterResponse {
    message: string;
    status: string;
}