"use client";

import { use } from "react";
import useSWR from "swr";
import Link from "next/link";
import { ArrowLeft, ExternalLink, Globe, Lock, ShieldCheck, User, Calendar } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardHeader, CardTitle, CardContent, CardDescription, CardFooter } from "@/components/ui/card";
import { mcpServerService } from "@/services/mcpServerService";
import { useAuthStore } from "@/store/useAuthStore";

export default function MCPServerDetailPage({ params }: { params: Promise<{ id: string }> }) {
  const { id: idStr } = use(params);
  const id = parseInt(idStr);
  const { user } = useAuthStore();

  const { data: serverResult, isLoading } = useSWR(
    id ? `/api/mcpservers/${id}` : null,
    () => mcpServerService.getMCPServer(id)
  );

  const server = serverResult?.data;

  if (isLoading) return <div className="p-10 text-center">Loading...</div>;
  if (!server) return <div className="p-10 text-center">Server not found</div>;

  const isOwner = user && server.creatorId === user.id;

  return (
    <div className="container mx-auto py-10 max-w-3xl px-4">
      <div className="mb-6 flex justify-between items-center">
        <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
            <ArrowLeft className="h-4 w-4 mr-1" /> 返回列表
        </Link>
        {isOwner && (
            <Link href={`/mcpservers/${id}/edit`}>
                <Button>编辑服务</Button>
            </Link>
        )}
      </div>

      <Card className="overflow-hidden">
        <CardHeader className="bg-muted/10 border-b">
           <div className="flex justify-between items-start">
              <div>
                  <CardTitle className="text-2xl mb-2">{server.name}</CardTitle>
                  <div className="flex flex-wrap gap-2">
                       <span className={`text-xs px-2 py-1 rounded-full border ${server.isPublic ? 'bg-green-50 text-green-700 border-green-200' : 'bg-yellow-50 text-yellow-700 border-yellow-200'} flex items-center gap-1`}>
                            {server.isPublic ? <Globe className="h-3 w-3" /> : <Lock className="h-3 w-3" />}
                            {server.isPublic ? '公开' : '私有'}
                       </span>
                       {server.openProxy && (
                            <span className="text-xs px-2 py-1 rounded-full border bg-blue-50 text-blue-700 border-blue-200 flex items-center gap-1">
                                <ShieldCheck className="h-3 w-3" /> Proxy Enabled
                            </span>
                        )}
                  </div>
              </div>
           </div>
        </CardHeader>
        <CardContent className="p-6 space-y-6">
            <div>
                <h3 className="text-sm font-medium text-muted-foreground mb-2">服务地址 (URL)</h3>
                <div className="bg-muted/30 p-3 rounded font-mono text-sm break-all flex items-center justify-between">
                    {server.url}
                    <a href={server.url} target="_blank" rel="noopener noreferrer" className="text-primary hover:underline ml-2">
                        <ExternalLink className="h-4 w-4" />
                    </a>
                </div>
            </div>

            <div>
                <h3 className="text-sm font-medium text-muted-foreground mb-2">描述</h3>
                <p className="whitespace-pre-wrap text-sm leading-relaxed">
                    {server.description || "暂无描述"}
                </p>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-4 border-t">
                 <div className="flex items-center gap-2 text-sm text-muted-foreground">
                     <User className="h-4 w-4" />
                     <span>创建者: {server.creatorNickname}</span>
                 </div>
                 <div className="flex items-center gap-2 text-sm text-muted-foreground">
                     <Calendar className="h-4 w-4" />
                     <span>创建时间: {server.createdAt ? new Date(server.createdAt).toLocaleDateString() : 'N/A'}</span>
                 </div>
            </div>
        </CardContent>
      </Card>
    </div>
  );
}