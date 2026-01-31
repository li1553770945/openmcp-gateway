"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
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
import { authService } from "@/services/authService";
import { useAuthStore } from "@/store/useAuthStore";
import { userService } from "@/services/userService";

export default function LoginPage() {
  const router = useRouter();
  const setToken = useAuthStore((state) => state.setToken);
  const setUser = useAuthStore((state) => state.setUser);
  
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [formData, setFormData] = useState({
    username: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      const resp = await authService.login(formData);
      if (resp.code === 0 && resp.data.token) {
        setToken(resp.data.token);
        
        // Fetch user info immediately after login to populate store
        try {
          // We need to wait a bit for the store to update or just use the token in the next request 
          // The fetchClient reads from store.getState().token, so it should be fine if we set it synchronously.
          
          // Use setTimeout to ensure state update propagation if needed, but getState is direct access.
          // However, react state updates are batched. Zustand updates outside react loop are usually sync.
          // Let's explicitly pass token if needed but our fetchClient handles it.
          // Small race condition might happen if zustand hasn't updated its internal state before fetchClient reads it?
          // Default create store is sync. 

          // Just call it.
          const userResp = await userService.getMe();
          if (userResp.code === 0) {
            setUser(userResp.data);
          }
        } catch (err) {
          console.error("Failed to fetch user info", err);
        }

        router.push("/");
      } else {
        setError(resp.message || "登录失败");
      }
    } catch (err: any) {
      setError(err.message || "登录发生错误");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto flex h-screen w-screen flex-col items-center justify-center">
      <Card className="w-[350px]">
        <CardHeader>
          <CardTitle>登录</CardTitle>
          <CardDescription>
            输入您的账号密码以登录 OpenMCP 网关
          </CardDescription>
        </CardHeader>
        <form onSubmit={handleSubmit}>
          <CardContent>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="username">用户名</Label>
                <Input
                  id="username"
                  name="username"
                  placeholder="请输入用户名"
                  value={formData.username}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="password">密码</Label>
                <Input
                  id="password"
                  name="password"
                  type="password"
                  placeholder="请输入密码"
                  value={formData.password}
                  onChange={handleChange}
                  required
                />
              </div>
              {error && (
                <div className="text-red-500 text-sm">{error}</div>
              )}
            </div>
          </CardContent>
          <CardFooter className="flex flex-col gap-4">
            <Button className="w-full" type="submit" disabled={loading}>
              {loading ? "登录中..." : "登录"}
            </Button>
            <div className="text-sm text-center text-muted-foreground">
              还没有账号?{" "}
              <Link href="/auth/register" className="text-primary hover:underline">
                去注册
              </Link>
            </div>
          </CardFooter>
        </form>
      </Card>
    </div>
  );
}
