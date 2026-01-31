"use client";

import Link from "next/link";
import useSWR from "swr";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { mcpServerService } from "@/services/mcpServerService";
import { useAuthStore } from "@/store/useAuthStore";
import { toast } from "sonner";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { 
  Globe, 
  Lock, 
  Edit, 
  Trash2, 
  Plus, 
  Server,
  ChevronLeft,
  ChevronRight,
  ShieldCheck,
  User,
} from "lucide-react";

export default function Home() {
  const router = useRouter();
  const { user } = useAuthStore();
  const [scope, setScope] = useState<'public' | 'self'>('public');
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(9);

  const start = (page - 1) * pageSize;
  const end = page * pageSize;

  const { data, error, isLoading, mutate } = useSWR(
    [`/api/mcpservers`, scope, page, pageSize],
    () => mcpServerService.getMCPServers({ scope, start, end })
  );

  const { data: countData } = useSWR(
    [`/api/mcpservers/count`, scope],
    () => mcpServerService.getMCPServerCount({ scope })
  );

  const totalCount = countData?.data?.count || 0;
  const totalPages = Math.max(1, Math.ceil(totalCount / pageSize));

  const handleDelete = async (id: number) => {
    if (confirm("确定要删除这个服务吗？")) {
       try {
         await mcpServerService.deleteMCPServer(id);
         toast.success("删除成功");
         mutate();
       } catch (e) {
         toast.error("删除失败");
       }
    }
  };

  const handleScopeChange = (newScope: 'public' | 'self') => {
    if (newScope === 'self' && !user) {
        toast.error("请先登录查看我的服务");
        router.push("/auth/login");
        return;
    }
    setScope(newScope);
    setPage(1);
  };

  const servers = data?.data || [];


  return (
    <div className="container mx-auto py-10 px-4 min-h-screen flex flex-col">
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center mb-8 gap-4">
        <div>
           <h1 className="text-3xl font-bold tracking-tight mb-2">MCPServer 列表</h1>
           <p className="text-muted-foreground">管理和发现 MCP 服务</p>
        </div>
        
        <div className="flex items-center gap-4">
             <div className="bg-muted p-1 rounded-lg flex text-sm font-medium">
                <button 
                    onClick={() => handleScopeChange('public')}
                    className={`px-4 py-2 rounded-md transition-all ${scope === 'public' ? 'bg-background shadow-sm text-foreground' : 'text-muted-foreground hover:text-foreground'}`}
                >
                    公共服务
                </button>
                <button 
                    onClick={() => handleScopeChange('self')}
                    className={`px-4 py-2 rounded-md transition-all ${scope === 'self' ? 'bg-background shadow-sm text-foreground' : 'text-muted-foreground hover:text-foreground'}`}
                >
                    我的服务
                </button>
             </div>
             <select 
                className="h-10 rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
                value={pageSize}
                onChange={(e) => {
                    setPageSize(Number(e.target.value));
                    setPage(1);
                }}
             >
                 <option value={9}>9 / 页</option>
                 <option value={20}>20 / 页</option>
                 <option value={50}>50 / 页</option>
             </select>
             <Link href="/mcpservers/create">
                <Button className="gap-2">
                    <Plus className="h-4 w-4" /> 添加服务
                </Button>
            </Link>
        </div>
      </div>

      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
             {[1, 2, 3, 4, 5, 6].map((i) => (
               <div key={i} className="h-[250px] rounded-lg border bg-card text-card-foreground shadow-sm animate-pulse bg-muted" />
             ))}
        </div>
      ) : servers.length > 0 ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
          {servers.map((server) => {
            const isOwner = user && server.creatorId === user.id;
            return (
                <Card key={server.id} className="flex flex-col h-full hover:shadow-md transition-shadow">
                <CardHeader>
                    <div className="flex justify-between items-start gap-2">
                         <div className="flex items-center gap-2">
                             <div className="p-2 bg-primary/10 rounded-full text-primary">
                                 <Server className="h-5 w-5" />
                             </div>
                             <div>
                                 <CardTitle className="text-lg line-clamp-1" title={server.name}>{server.name}</CardTitle>
                                 <div className="flex items-center gap-2 mt-1">
                                    <span className={`text-[10px] px-2 py-0.5 rounded-full border ${server.isPublic ? 'bg-green-50 text-green-700 border-green-200' : 'bg-yellow-50 text-yellow-700 border-yellow-200'} flex items-center gap-1`}>
                                        {server.isPublic ? <Globe className="h-3 w-3" /> : <Lock className="h-3 w-3" />}
                                        {server.isPublic ? '公开' : '私有'}
                                    </span>
                                    {server.openProxy && (
                                        <span className="text-[10px] px-2 py-0.5 rounded-full border bg-blue-50 text-blue-700 border-blue-200 flex items-center gap-1">
                                            <ShieldCheck className="h-3 w-3" /> Proxy
                                        </span>
                                    )}
                                 </div>
                             </div>
                         </div>
                    </div>
                </CardHeader>
                <CardContent className="flex-1">
                    <div className="text-sm text-muted-foreground mb-4 break-all bg-muted/50 p-2 rounded text-xs font-mono">
                         {server.url}
                    </div>
                    <p className="text-sm text-muted-foreground line-clamp-3">
                    {server.description || "暂无描述"}
                    </p>
                    <div className="mt-4 flex items-center gap-1 text-xs text-muted-foreground">
                        <User className="h-3 w-3" />
                        <span>{server.creatorNickname}</span>
                    </div>
                </CardContent>
                <CardFooter className="pt-2 border-t flex justify-between bg-muted/10">
                    <div className="text-xs text-muted-foreground font-mono opacity-50">
                        #{server.id}
                    </div>
                    <div className="flex gap-2">
                        {isOwner && (
                            <Button variant="ghost" size="icon" className="h-8 w-8 text-destructive hover:text-destructive hover:bg-destructive/10" onClick={() => handleDelete(server.id)} title="删除">
                                <Trash2 className="h-4 w-4" />
                            </Button>
                        )}
                         {isOwner && (
                             <Link href={`/mcpservers/${server.id}/edit`}>
                                <Button variant="outline" size="sm" className="h-8 gap-1">
                                    <Edit className="h-3 w-3" /> 编辑
                                </Button>
                            </Link>
                         )}
                         <Link href={`/mcpservers/${server.id}`}>
                            <Button variant="secondary" size="sm" className="h-8">
                                详情
                            </Button>
                         </Link>
                    </div>
                </CardFooter>
                </Card>
            );
          })}
        </div>
      ) : (
          <div className="flex flex-col items-center justify-center py-20 text-center border-2 border-dashed rounded-lg bg-muted/10">
              <Server className="h-10 w-10 text-muted-foreground mb-4" />
              <h3 className="text-lg font-medium">没有找到服务</h3>
              <p className="text-muted-foreground mb-4">
                  {scope === 'self' ? '你还没有创建任何服务' : '暂时没有公开的服务'}
              </p>
              {scope === 'self' && (
                 <Link href="/mcpservers/create">
                    <Button>创建服务</Button>
                 </Link>
              )}
          </div>
      )}

      {/* Pagination */}
      <div className="mt-auto flex flex-col sm:flex-row justify-between items-center gap-4 py-4 pt-8 border-t">
          <div className="text-sm text-muted-foreground">
              共 {totalCount} 个服务，第 {page} / {totalPages} 页
          </div>
          <div className="flex items-center gap-2">
            <Button 
                variant="outline" 
                size="icon" 
                onClick={() => setPage(p => Math.max(1, p - 1))}
                disabled={page === 1 || isLoading}
            >
                <ChevronLeft className="h-4 w-4" />
            </Button>
            <Button 
                variant="outline" 
                size="icon" 
                onClick={() => setPage(p => Math.min(totalPages, p + 1))}
                disabled={page >= totalPages || isLoading}
            >
                <ChevronRight className="h-4 w-4" />
            </Button>
          </div>
      </div>
    </div>
  );
}
