"use client";

import { useEffect, useState, use } from "react";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { Switch } from "@/components/ui/switch";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog";
import { Card, CardContent, CardFooter, CardHeader, CardTitle, CardDescription } from "@/components/ui/card";
import { mcpServerService } from "@/services/mcpServerService";
import { UpdateMCPServerReq } from "@/types/mcpserver";
import { ArrowLeft, Trash2, Copy } from "lucide-react";
import Link from "next/link";
import useSWR from "swr";

export default function EditMCPServerPage({ params }: { params: Promise<{ id: string }> }) {
  const { id: idStr } = use(params);
  const id = parseInt(idStr);
  const router = useRouter();

  const { data: serverResult, mutate, isLoading } = useSWR(
    id ? `/api/mcpservers/${id}` : null,
    () => mcpServerService.getMCPServer(id)
  );

  const server = serverResult?.data;

  const [formData, setFormData] = useState<UpdateMCPServerReq>({
    id: id,
    name: "",
    url: "",
    description: "",
    isPublic: false,
    openProxy: false,
  });

  const [saving, setSaving] = useState(false);
  const [newTokenDesc, setNewTokenDesc] = useState("");
  const [generating, setGenerating] = useState(false);
  const [lastGeneratedToken, setLastGeneratedToken] = useState<string | null>(null);
  const [tokenToDelete, setTokenToDelete] = useState<number | null>(null);

  useEffect(() => {
    if (server) {
      setFormData({
        id: server.id,
        name: server.name,
        url: server.url,
        description: server.description || "",
        isPublic: server.isPublic,
        openProxy: server.openProxy,
      });
    }
  }, [server]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleUpdate = async (e: React.FormEvent) => {
    e.preventDefault();
    setSaving(true);
    try {
      await mcpServerService.updateMCPServer(formData);
      toast.success("更新成功");
      mutate();
    } catch (error) {
      toast.error("更新失败", {
        description: error instanceof Error ? error.message : "发生未知错误"
      });
      console.error(error);
    } finally {
      setSaving(false);
    }
  };

  const handleGenerateToken = async () => {
    if (!newTokenDesc.trim()) {
        toast.error("请输入 Token 描述");
        return;
    }
    setGenerating(true);
    try {
        const res = await mcpServerService.generateToken({
            id: id,
            description: newTokenDesc
        });
        if (res.data?.token) {
            setLastGeneratedToken(res.data.token);
            setNewTokenDesc("");
            toast.success("Token 生成成功！");
            mutate(); // Refresh token list
        }
    } catch (error) {
        toast.error("Generate Token Failed", {
            description: error instanceof Error ? error.message : "发生未知错误"
        });
    } finally {
        setGenerating(false);
    }
  };

  const handleDeleteToken = async () => {
    if (!tokenToDelete) return;
    try {
        await mcpServerService.deleteToken(tokenToDelete);
        toast.success("Token 删除成功");
        mutate();
    } catch (e) {
        toast.error("删除 Token 失败", {
            description: e instanceof Error ? e.message : "发生未知错误"
        });
    } finally {
        setTokenToDelete(null);
    }
  };

  const copyToClipboard = (text: string) => {
      navigator.clipboard.writeText(text);
      toast.success("已复制到剪贴板");
  };

  if (isLoading) return <div className="p-10 text-center">Loading...</div>;
  if (!server) return <div className="p-10 text-center">Server not found</div>;

  return (
    <div className="container mx-auto py-10 max-w-4xl px-4">
        <div className="mb-6">
            <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
                <ArrowLeft className="h-4 w-4 mr-1" /> 返回列表
            </Link>
        </div>

      <div className="grid gap-6">
          {/* Basic Info */}
          <Card>
            <CardHeader>
            <CardTitle>基本信息</CardTitle>
            <CardDescription>编辑服务的基本配置</CardDescription>
            </CardHeader>
            <form onSubmit={handleUpdate}>
            <CardContent className="space-y-4">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                    <Label htmlFor="name">名称</Label>
                    <Input
                        id="name"
                        name="name"
                        required
                        value={formData.name}
                        onChange={handleChange}
                    />
                    </div>
                    <div className="space-y-2">
                    <Label htmlFor="url">URL</Label>
                    <Input
                        id="url"
                        name="url"
                        required
                        type="url"
                        value={formData.url}
                        onChange={handleChange}
                    />
                    </div>
                </div>

                <div className="space-y-2">
                <Label htmlFor="description">描述</Label>
                <Textarea
                    id="description"
                    name="description"
                    className="min-h-[80px]"
                    value={formData.description}
                    onChange={handleChange}
                />
                </div>

                <div className="flex flex-col gap-4 sm:flex-row sm:items-center">
                     <div className="flex items-center justify-between border p-3 rounded-md flex-1">
                        <div className="flex flex-col space-y-0.5">
                            <Label htmlFor="isPublic" className="cursor-pointer">公开服务</Label>
                            <span className="text-xs text-muted-foreground">允许其他人查看此服务</span>
                        </div>
                        <Switch
                            id="isPublic"
                            checked={formData.isPublic}
                            onCheckedChange={(checked) => setFormData(prev => ({ ...prev, isPublic: checked }))}
                        />
                    </div>

                    <div className="flex items-center justify-between border p-3 rounded-md flex-1">
                         <div className="flex flex-col space-y-0.5">
                            <Label htmlFor="openProxy" className="cursor-pointer">开启代理</Label>
                             <span className="text-xs text-muted-foreground">允许通过网关代理请求</span>
                        </div>
                        <Switch
                            id="openProxy"
                            checked={formData.openProxy}
                            onCheckedChange={(checked) => setFormData(prev => ({ ...prev, openProxy: checked }))}
                        />
                    </div>
                </div>

            </CardContent>
            <CardFooter className="flex justify-end gap-2 bg-muted/20 border-t p-4">
                <Button type="submit" disabled={saving}>
                    {saving ? "保存中..." : "保存更改"}
                </Button>
            </CardFooter>
            </form>
          </Card>

          {/* Tokens */}
          <Card>
              <CardHeader>
                  <CardTitle>访问令牌 (Tokens)</CardTitle>
                  <CardDescription>管理用于访问此服务的 Token</CardDescription>
              </CardHeader>
              <CardContent>
                  {/* Generate Token Section */}
                  <div className="mb-6 p-4 border rounded-md bg-muted/30">
                      <h4 className="text-sm font-medium mb-3">生成新 Token</h4>
                      <div className="flex gap-2">
                          <Input 
                              placeholder="Token 描述 (例如: Web Client)" 
                              value={newTokenDesc}
                              onChange={(e) => setNewTokenDesc(e.target.value)}
                          />
                          <Button onClick={handleGenerateToken} disabled={generating}>
                              {generating ? "生成中..." : "生成"}
                          </Button>
                      </div>
                      {lastGeneratedToken && (
                          <div className="mt-4 p-3 bg-green-50 border border-green-200 rounded-md text-green-800">
                              <p className="text-xs font-bold mb-1">新生成的 Token:</p>
                              <div className="flex items-center justify-between gap-2">
                                  <code className="bg-white px-2 py-1 rounded border border-green-200 text-sm break-all flex-1 select-all">
                                      {lastGeneratedToken}
                                  </code>
                                  <Button variant="ghost" size="icon" className="h-8 w-8 text-green-700 hover:bg-green-100" onClick={() => copyToClipboard(lastGeneratedToken)}>
                                      <Copy className="h-4 w-4" />
                                  </Button>
                              </div>
                          </div>
                      )}
                  </div>

                  {/* Token List */}
                  <div className="space-y-4">
                      <h4 className="text-sm font-medium">现有 Token</h4>
                      {server.token && server.token.length > 0 ? (
                          <div className="space-y-2">
                              {server.token.map(t => (
                                  <div key={t.id} className="flex items-center justify-between p-3 border rounded bg-card hover:bg-muted/10 transition-colors">
                                      <div className="flex-1 min-w-0 mr-4">
                                          <div className="font-medium text-sm truncate">{t.description}</div>
                                          <div className="text-xs text-muted-foreground font-mono mt-1 truncate">
                                              {t.token ? t.token.substring(0, 10) + '...' : '****************'}
                                          </div>
                                      </div>
                                      <div className="flex items-center gap-2">
                                           {t.token && (
                                                <Button variant="ghost" size="icon" className="h-8 w-8" onClick={() => copyToClipboard(t.token)} title="复制完整 Token">
                                                    <Copy className="h-4 w-4" />
                                                </Button>
                                           )}
                                          <Button variant="ghost" size="icon" className="h-8 w-8 text-destructive hover:text-destructive hover:bg-destructive/10" onClick={() => setTokenToDelete(t.id)} title="删除">
                                              <Trash2 className="h-4 w-4" />
                                          </Button>
                                      </div>
                                  </div>
                              ))}
                          </div>
                      ) : (
                          <div className="text-sm text-muted-foreground text-center py-4">暂无 Token</div>
                      )}
                  </div>
              </CardContent>
          </Card>
      </div>
      
      <AlertDialog open={!!tokenToDelete} onOpenChange={(open) => !open && setTokenToDelete(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>确认删除</AlertDialogTitle>
            <AlertDialogDescription>
              您确定要删除这个 Token 吗？此操作无法撤销。
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>取消</AlertDialogCancel>
            <AlertDialogAction onClick={handleDeleteToken} className="bg-destructive text-destructive-foreground hover:bg-destructive/90">删除</AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  );
}
