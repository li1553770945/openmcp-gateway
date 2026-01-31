import { fetchClient } from '@/lib/api-client';
import { 
  GetMCPServerListResp, 
  GetMCPServerListReq,
  GetMCPServerCountReq,
  GetMCPServerCountResp,
  AddMCPServerReq, 
  AddMCPServerResp, 
  UpdateMCPServerReq, 
  UpdateMCPServerResp, 
  DeleteMCPServerResp,
  GetMCPServerResp,
  GenerateTokenReq,
  GenerateTokenResp,
  DeleteTokenResp
} from '@/types/mcpserver';

export const mcpServerService = {
  async getMCPServers(params: GetMCPServerListReq): Promise<GetMCPServerListResp> {
    const searchParams = new URLSearchParams();
    if (params.scope) searchParams.append('scope', params.scope);
    if (params.start !== undefined) searchParams.append('start', params.start.toString());
    if (params.end !== undefined) searchParams.append('end', params.end.toString());
    
    return fetchClient<GetMCPServerListResp>(`/api/mcpservers?${searchParams.toString()}`);
  },

  async getMCPServerCount(params: GetMCPServerCountReq): Promise<GetMCPServerCountResp> {
    const searchParams = new URLSearchParams();
    if (params.scope) searchParams.append('scope', params.scope);
    
    return fetchClient<GetMCPServerCountResp>(`/api/mcpservers/count?${searchParams.toString()}`);
  },

  async getMCPServer(id: number): Promise<GetMCPServerResp> {
    return fetchClient<GetMCPServerResp>(`/api/mcpservers/${id}`);
  },

  async addMCPServer(data: AddMCPServerReq): Promise<AddMCPServerResp> {
    return fetchClient<AddMCPServerResp>('/api/mcpservers', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  },

  async updateMCPServer(data: UpdateMCPServerReq): Promise<UpdateMCPServerResp> {
    return fetchClient<UpdateMCPServerResp>(`/api/mcpservers/${data.id}`, {
      method: 'PUT',
      body: JSON.stringify(data),
    });
  },

  async deleteMCPServer(id: number): Promise<DeleteMCPServerResp> {
    return fetchClient<DeleteMCPServerResp>(`/api/mcpservers/${id}`, {
      method: 'DELETE',
    });
  },

  async generateToken(data: GenerateTokenReq): Promise<GenerateTokenResp> {
    return fetchClient<GenerateTokenResp>('/api/mcpservers/tokens', {
      method: 'POST',
      body: JSON.stringify(data),
    });
  },

  async deleteToken(id: number): Promise<DeleteTokenResp> {
    return fetchClient<DeleteTokenResp>(`/api/mcpservers/tokens/${id}`, {
      method: 'DELETE',
    });
  }
};
