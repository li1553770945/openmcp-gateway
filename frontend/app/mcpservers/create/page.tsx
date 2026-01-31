"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { mcpServerService } from "@/services/mcpServerService";
import { AddMCPServerReq } from "@/types/mcpserver";
import { ArrowLeft } from "lucide-react";
import Link from "next/link";

export default function CreateMCPServerPage() {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [formData, setFormData] = useState<AddMCPServerReq>({
    name: "",
    url: "",
    description: "",
    isPublic: false,
    openProxy: false,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    const { name, value, type } = e.target;
    if (type === 'checkbox') {
        const checked = (e.target as HTMLInputElement).checked;
        setFormData(prev => ({ ...prev, [name]: checked }));
    } else {
        setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    try {
      await mcpServerService.addMCPServer(formData);
      toast.success("创建成功");
      router.push("/");
    } catch (error) {
      toast.error("创建失败", {
        description: error instanceof Error ? error.message : "发生未知错误"
      });
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto py-10 max-w-2xl px-4">
        <div className="mb-6">
            <Link href="/" className="flex items-center text-sm text-muted-foreground hover:text-foreground">
                <ArrowLeft className="h-4 w-4 mr-1" /> 返回列表
            </Link>
        </div>
      <Card>
        <CardHeader>
          <CardTitle>创建新的 MCP 服务</CardTitle>
        </CardHeader>
        <form onSubmit={handleSubmit}>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="name">名称</Label>
              <Input
                id="name"
                name="name"
                required
                value={formData.name}
                onChange={handleChange}
                placeholder="给你的服务起个名字"
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
                placeholder="http://example.com/mcp"
              />
            </div>

            <div className="space-y-2">
              <Label htmlFor="description">描述</Label>
              <textarea
                id="description"
                name="description"
                className="flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                value={formData.description}
                onChange={handleChange}
                placeholder="描述这个服务的功能..."
              />
            </div>

            <div className="flex items-center space-x-2">
               <input
                type="checkbox"
                id="isPublic"
                name="isPublic"
                checked={formData.isPublic}
                onChange={handleChange}
                className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary"
               />
               <Label htmlFor="isPublic">公开服务 (允许其他人查看)</Label>
            </div>

            <div className="flex items-center space-x-2">
               <input
                type="checkbox"
                id="openProxy"
                name="openProxy"
                checked={formData.openProxy}
                onChange={handleChange}
                className="h-4 w-4 rounded border-gray-300 text-primary focus:ring-primary"
               />
               <Label htmlFor="openProxy">开启代理 (允许通过网关代理请求)</Label>
            </div>

          </CardContent>
          <CardFooter className="flex justify-end gap-2">
            <Button variant="outline" type="button" onClick={() => router.back()}>取消</Button>
            <Button type="submit" disabled={loading}>
                {loading ? "创建中..." : "创建"}
            </Button>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
