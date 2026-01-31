"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import useSWR, { mutate } from "swr";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { userService } from "@/services/userService";
import { useAuthStore } from "@/store/useAuthStore";
import { notificationSuccess, notificationError } from "@/lib/notification";
import { Loader2, User as UserIcon, LogOut, Edit2, Save, X, AlertCircle } from "lucide-react";

const fetcher = () => userService.getMe().then((res) => res.data);

export default function UserProfilePage() {
  const router = useRouter();
  const logout = useAuthStore((state) => state.logout);
  const setUser = useAuthStore((state) => state.setUser);
  const { data: user, error, isLoading } = useSWR("user-me", fetcher);
  
  const [isEditing, setIsEditing] = useState(false);
  const [formData, setFormData] = useState({
    nickname: "",
    password: "", 
  });
  const [updateLoading, setUpdateLoading] = useState(false);

  useEffect(() => {
    if (user) {
      setFormData((prev) => ({
        ...prev,
        nickname: user.nickname || "",
      }));
      setUser(user);
    }
  }, [user, setUser]);

  const handleLogout = () => {
    logout();
    notificationSuccess("已退出登录");
    router.push("/auth/login");
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setUpdateLoading(true);

    try {
      const payload: any = { nickname: formData.nickname };
      if (formData.password) {
        payload.password = formData.password;
      }
      
      const resp = await userService.updateMe(payload);
      if (resp.code === 0) {
        mutate("user-me"); // Revalidate SWR
        setIsEditing(false);
        setFormData(prev => ({ ...prev, password: "" })); // Clear password
        notificationSuccess("更新成功", "您的个人信息已成功修改");
      } else {
        notificationError("更新失败", resp.message || "请稍后重试");
      }
    } catch (err: any) {
      notificationError("更新发生错误", err.message);
    } finally {
      setUpdateLoading(false);
    }
  };

  if (isLoading) {
    return (
        <div className="flex h-[60vh] w-full flex-col items-center justify-center gap-4">
            <Loader2 className="h-10 w-10 animate-spin text-primary" />
            <p className="text-muted-foreground animate-pulse">正在加载用户信息...</p>
        </div>
    );
  }

  if (error) {
    return (
         <div className="flex h-[60vh] w-full items-center justify-center p-4">
             <Card className="w-full max-w-md shadow-lg border-destructive/20">
                <CardHeader className="text-center">
                    <div className="mx-auto mb-4 flex h-12 w-12 items-center justify-center rounded-full bg-destructive/10">
                        <AlertCircle className="h-6 w-6 text-destructive" />
                    </div>
                    <CardTitle className="text-xl text-destructive">无法获取用户信息</CardTitle>
                    <CardDescription>您的登录状态可能已失效，请重新登录。</CardDescription>
                </CardHeader>
                <CardFooter className="flex justify-center pb-6">
                     <Button variant="default" onClick={() => router.push('/auth/login')} className="w-full">
                        返回登录页面
                     </Button>
                </CardFooter>
             </Card>
        </div>
    );
  }

  return (
    <div className="container mx-auto max-w-4xl px-4 py-8 md:py-12">
        <div className="mb-8 space-y-2">
             <h1 className="text-3xl font-bold tracking-tight">个人中心</h1>
             <p className="text-muted-foreground">
                管理您的个人资料和账号安全设置
             </p>
        </div>
        
        <Card className="overflow-hidden border shadow-sm">
            <CardHeader className="border-b bg-muted/40 px-6 py-8">
                <div className="flex flex-col gap-6 md:flex-row md:items-center md:justify-between">
                    <div className="flex items-center gap-5">
                        <div className="flex h-20 w-20 shrink-0 items-center justify-center rounded-full bg-primary/10 shadow-inner">
                            <UserIcon className="h-10 w-10 text-primary" />
                        </div>
                        <div className="space-y-1.5">
                            <CardTitle className="text-2xl">{user?.nickname || user?.username}</CardTitle>
                            <div className="flex items-center gap-2 text-sm text-muted-foreground">
                                <span className={`inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium ${
                                    user?.role === 'admin' 
                                        ? 'bg-primary/10 text-primary' 
                                        : 'bg-secondary text-secondary-foreground'
                                }`}>
                                    {user?.role === 'admin' ? '管理员' : '普通用户'}
                                </span>
                                <span>ID: {user?.id}</span>
                            </div>
                        </div>
                    </div>
                    {!isEditing && (
                        <div className="flex gap-3">
                             <Button variant="outline" onClick={() => setIsEditing(true)}>
                                <Edit2 className="mr-2 h-4 w-4" />
                                编辑资料
                            </Button>
                        </div>
                    )}
                </div>
            </CardHeader>
            <CardContent className="p-6 md:p-8">
                {isEditing ? (
                    <form id="profile-form" onSubmit={handleSubmit} className="mx-auto max-w-2xl space-y-6">
                        <div className="grid gap-6 md:grid-cols-2">
                             <div className="space-y-2">
                                <Label htmlFor="username">用户名</Label>
                                <Input 
                                    id="username" 
                                    value={user?.username} 
                                    disabled 
                                    className="bg-muted text-muted-foreground cursor-not-allowed"
                                />
                                <p className="text-[0.8rem] text-muted-foreground">用户名不可修改</p>
                            </div>
                            <div className="space-y-2">
                                <Label htmlFor="nickname">昵称</Label>
                                <Input 
                                    id="nickname" 
                                    name="nickname" 
                                    value={formData.nickname} 
                                    onChange={handleChange} 
                                    placeholder="您的昵称"
                                    className="bg-background"
                                />
                            </div>
                        </div>
                         <div className="space-y-3 pt-2">
                            <Label htmlFor="password">新密码</Label>
                            <Input 
                                id="password" 
                                name="password" 
                                type="password"
                                placeholder="请输入新密码（如不修改请留空）"
                                value={formData.password} 
                                onChange={handleChange} 
                                className="bg-background"
                            />
                             <p className="text-[0.8rem] text-muted-foreground flex items-center gap-1">
                                <AlertCircle className="h-3 w-3" />
                                留空则保持原密码不变
                            </p>
                        </div>
                    </form>
                ) : (
                    <dl className="grid gap-x-8 gap-y-6 md:grid-cols-2 lg:grid-cols-3">
                         <div className="space-y-1.5">
                            <dt className="text-sm font-medium text-muted-foreground">用户 ID</dt>
                            <dd className="font-medium text-foreground">{user?.id}</dd>
                         </div>
                         <div className="space-y-1.5">
                            <dt className="text-sm font-medium text-muted-foreground">用户名</dt>
                            <dd className="font-medium text-foreground">{user?.username}</dd>
                         </div>
                         <div className="space-y-1.5">
                            <dt className="text-sm font-medium text-muted-foreground">昵称</dt>
                            <dd className="font-medium text-foreground">{user?.nickname || "未设置"}</dd>
                         </div>
                         <div className="space-y-1.5">
                            <dt className="text-sm font-medium text-muted-foreground">角色权限</dt>
                            <dd className="font-medium text-foreground">{user?.role === 'admin' ? '系统管理员' : '普通注册用户'}</dd>
                         </div>
                    </dl>
                )}
            </CardContent>
            <CardFooter className="flex items-center justify-between border-t bg-muted/20 px-6 py-4">
                 {isEditing ? (
                    <div className="flex w-full justify-end gap-3">
                        <Button type="button" variant="ghost" onClick={() => setIsEditing(false)} disabled={updateLoading}>
                            <X className="mr-2 h-4 w-4" />
                            取消
                        </Button>
                        <Button type="submit" form="profile-form" disabled={updateLoading} className="min-w-[120px]">
                            {updateLoading ? (
                                <>
                                    <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                                    保存中...
                                </>
                            ) : (
                                <>
                                    <Save className="mr-2 h-4 w-4" />
                                    保存更改
                                </>
                            )}
                        </Button>
                    </div>
                 ) : (
                    <div className="flex w-full items-center justify-between">
                         <div />
                         <Button variant="destructive" size="sm" onClick={handleLogout} className="gap-2">
                            <LogOut className="h-4 w-4" />
                            退出登录
                        </Button>
                    </div>
                 )}
            </CardFooter>
        </Card>
    </div>
  );
}
