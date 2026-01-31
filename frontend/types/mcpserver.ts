import { BaseResponse } from './common';

// MCPServer管理
export interface AddMCPServerReq {
  name: string;
  description?: string;
  url: string;
  isPublic?: boolean;
  openProxy?: boolean;
}

export type AddMCPServerResp = BaseResponse<null>;

export interface UpdateMCPServerReq {
  id: number;
  name?: string;
  description?: string;
  url?: string;
  isPublic?: boolean;
  openProxy?: boolean;
}

export type UpdateMCPServerResp = BaseResponse<null>;

export type DeleteMCPServerResp = BaseResponse<null>;

export interface MCPServerData {
  id: number;
  name: string;
  description: string;
  url: string;
  isPublic: boolean;
  openProxy: boolean;
  creatorId: number;
  creatorNickname: string;
  createdAt: string;
  updatedAt: string;
}

export interface TokenData {
  id: number;
  token: string;
  description: string;
}

export interface MCPServerDetailData extends MCPServerData {
  token?: TokenData[];
}

export interface GetMCPServerListReq {
  scope?: 'self' | 'public' ;
  start: number;
  end: number;
}

export type GetMCPServerListResp = BaseResponse<MCPServerData[]>;

export interface GetMCPServerCountReq {
  scope?: 'public' | 'self';
}

export interface GetMCPServerCountRespData {
  count: number;
}

export type GetMCPServerCountResp = BaseResponse<GetMCPServerCountRespData>;

export type GetMCPServerResp = BaseResponse<MCPServerDetailData>;

export interface GenerateTokenReq {
  id: number;
  description: string;
}

export interface GenerateTokenRespData {
  token: string;
}

export type GenerateTokenResp = BaseResponse<GenerateTokenRespData>;

export type DeleteTokenResp = BaseResponse<null>;
