import { fetchClient } from '@/lib/api-client';
import { GetUserInfoResp, UpdateSelfInfoReq, UpdateSelfInfoResp,RegisterReq,RegisterResp } from '@/types/user';

export const userService = {
    async getMe(): Promise<GetUserInfoResp> {
        return fetchClient<GetUserInfoResp>('/api/users/me');
    },

    async updateMe(data: UpdateSelfInfoReq): Promise<UpdateSelfInfoResp> {
        return fetchClient<UpdateSelfInfoResp>('/api/users/me', {
            method: 'PUT',
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
