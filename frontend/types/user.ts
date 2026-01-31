import { BaseResponse } from './common';

// 用户管理
export interface RegisterReq {
  username: string;
  password?: string;
  nickname?: string;
  email?: string;
}

export type RegisterResp = BaseResponse<null>;

export interface UpdateSelfInfoReq {
  nickname?: string;
  password?: string;
}

export type UpdateSelfInfoResp = BaseResponse<null>;

export interface UserInfo {
  id: number;
  username: string;
  nickname: string;
  role: string;
}

export type GetUserInfoResp = BaseResponse<UserInfo>;
