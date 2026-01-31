import { BaseResponse } from './common';

// 认证管理
export interface LoginReq {
  username: string;
  password?: string;
}

export interface LoginRespData {
  token: string;
}

export type LoginResp = BaseResponse<LoginRespData>;
