import { fetchClient } from '@/lib/api-client';
import { LoginReq, LoginResp } from '@/types/auth';
import { RegisterReq, RegisterResp } from '@/types/user';

export const authService = {
  async login(data: LoginReq): Promise<LoginResp> {
    return fetchClient<LoginResp>('/api/auth/login', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  },

  async register(data: RegisterReq): Promise<RegisterResp> {
    return fetchClient<RegisterResp>('/api/users', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  }
};
