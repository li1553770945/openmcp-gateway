"use client";

import Link from "next/link";
import { Button } from "@/components/ui/button";
import { useAuthStore } from "@/store/useAuthStore";
import { useEffect, useState } from "react";
import { User } from "lucide-react";

export function Navbar() {
  const user = useAuthStore((state) => state.user);
  const token = useAuthStore((state) => state.token);
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  return (
    <nav className="border-b bg-background">
      <div className="container mx-auto flex h-16 items-center px-4">
        <div className="mr-8">
          <Link href="/" className="text-xl font-bold">
            OpenMCP Gateway
          </Link>
        </div>
        <div className="flex gap-6 text-sm font-medium">
          <Link href="/" className="transition-colors hover:text-foreground/80">
            服务列表
          </Link>
          {mounted && token && (
             <Link
              href="/mcpservers/create"
              className="transition-colors hover:text-foreground/80"
            >
              创建服务
            </Link>
          )}
        </div>
        <div className="ml-auto flex items-center gap-4">
          {mounted ? (
            token ? (
              <>
                 <Link href="/users/me">
                  <Button variant="ghost" size="sm" className="gap-2">
                    <User className="h-4 w-4" />
                    {user?.nickname || user?.username || '用户'}
                  </Button>
                </Link>
              </>
            ) : (
              <>
                <Link href="/auth/login">
                  <Button variant="ghost" size="sm">
                    登录
                  </Button>
                </Link>
                <Link href="/auth/register">
                  <Button size="sm">注册</Button>
                </Link>
              </>
            )
          ) : (
            // Server/Hydration loading state - show nothing or skeleton to avoid flicker
            <div className="w-20" />
          )}
        </div>
      </div>
    </nav>
  );
}
