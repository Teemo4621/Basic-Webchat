'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Eye, EyeOff, User, Lock, ArrowRight } from 'lucide-react';
import { Label } from "@/components/ui/label";
import { Form, FormField, FormItem, FormControl, FormMessage } from "@/components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import { useForm } from "react-hook-form";
import useAuthentication from '@/hooks/authicationHook';

const loginSchema = z.object({
  username: z.string().min(3, { message: "Username must be at least 3 characters" }),
  password: z.string().min(6, { message: "Password must be at least 6 characters" }),
});

type LoginValues = z.infer<typeof loginSchema>;

export default function LoginPage() {
  const [showPassword, setShowPassword] = useState(false);
  const router = useRouter();
  const { login, success } = useAuthentication()

  useEffect(() => {
    if (success) {
      router.push('/rooms');
    }
  }, [success, router]);

  const form = useForm<LoginValues>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const onSubmit = (values: LoginValues) => {
    const data = {
      username: values.username,
      password: values.password,
    };
    
    login(data);
  };

  const toggleShowPassword = () => {
    setShowPassword(!showPassword);
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center p-4 py-12 relative overflow-hidden">
      
      <div className="glass p-8 rounded-2xl w-full max-w-md z-10 animate-fadeIn">
        <div className="space-y-6">
          <div className="space-y-2 text-center">
            <h1 className="text-3xl font-bold tracking-tighter">Welcome Back</h1>
            <p className="text-muted-foreground">Login to your account to continue</p>
          </div>

          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <div className="space-y-2">
                      <Label htmlFor="username">Username</Label>
                      <div className="relative">
                        <User className="absolute left-3 top-3 h-4 w-4 text-white/50" />
                        <FormControl>
                          <Input
                            id="username"
                            type="text"
                            placeholder="johndoe"
                            className="pl-10 bg-white/5 border-white/10 h-12"
                            autoComplete="off"
                            {...field}
                          />
                        </FormControl>
                      </div>
                      <FormMessage />
                    </div>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <div className="space-y-2">
                      <Label htmlFor="password">Password</Label>
                      <div className="relative">
                        <Lock className="absolute left-3 top-3 h-4 w-4 text-white/50" />
                        <FormControl>
                          <Input
                            id="password"
                            type={showPassword ? "text" : "password"}
                            placeholder="••••••••"
                            className="pl-10 bg-white/5 border-white/10 h-12"
                            {...field}
                          />
                        </FormControl>
                        <Button
                          type="button"
                          variant="ghost"
                          size="icon"
                          className="absolute right-2 top-3 h-6 w-6"
                          onClick={toggleShowPassword}
                        >
                          {showPassword ? (
                            <EyeOff className="h-4 w-4 text-white/50" />
                          ) : (
                            <Eye className="h-4 w-4 text-white/50" />
                          )}
                        </Button>
                      </div>
                      <FormMessage />
                    </div>
                  </FormItem>
                )}
              />

              <Button
                type="submit"
                className="w-full h-12 bg-white text-black hover:bg-white/90 transition-all"
              >
                Login
                <ArrowRight className="ml-2 h-4 w-4" />
              </Button>
            </form>
          </Form>

          <div className="text-center space-y-4">
            <p className="text-sm text-white/60">
              Don&apos;t have an account?{' '}
              <Link href="/register" className="text-white hover:underline">
                Register
              </Link>
            </p>
            <div className="flex justify-center">
              <Link href="/" className="text-sm text-white/60 hover:text-white">
                Back to Home
              </Link>
            </div>
          </div>
        </div>
      </div>

      <footer className="absolute bottom-4 text-xs text-white/40">
        Void Flow © {new Date().getFullYear()}
      </footer>
    </div>
  );
}