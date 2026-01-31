// 通用响应结构
export interface BaseResponse<T = any> {
  code: number;
  message: string;
  data: T;
}
