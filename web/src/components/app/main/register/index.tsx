"use client";

import { useRouter } from "next/navigation";
import Link from "next/link";
import { useMutation } from "@tanstack/react-query";
import { zodResolver } from "@hookform/resolvers/zod";
import { type SubmitHandler, useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "~/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "~/components/ui/form";
import { Input } from "~/components/ui/input";
import { useToast } from "~/components/ui/use-toast";

import api from "~/client/api";

const formSchema = z.object({
  email: z.string().email({ message: "Invalid email address" }),
  username: z
    .string()
    .min(3, { message: "Username must be at least 3 characters" })
    .max(20, { message: "Username must be at most 20 characters" }),
  password: z
    .string()
    .min(8, { message: "Password must be at least 8 characters" }),
});

type RegisterFormData = z.infer<typeof formSchema>;

export default function RegisterForm() {
  const router = useRouter();

  const { toast } = useToast();

  const form = useForm<RegisterFormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      username: "",
      password: "",
    },
    mode: "onBlur",
  });

  const { handleSubmit, reset, control } = form;

  const registerMutation = useMutation({
    mutationFn: async (data: RegisterFormData) =>
      await api.post("/users/register", data),
    onSuccess: () => {
      router.push("/login");
    },
    onError: (error) => {
      toast({
        variant: "destructive",
        title: "An error occurred",
        description: error.message,
      });

      reset();
    },
  });

  const { mutate, isPending } = registerMutation;

  const onSubmit: SubmitHandler<RegisterFormData> = (data) => mutate(data);

  return (
    <Form {...form}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Card className="w-[350px]">
          <CardHeader>
            <CardTitle>Register</CardTitle>
            <CardDescription>
              Create an account to start using the app
            </CardDescription>
          </CardHeader>

          <CardContent className="space-y-4">
            <FormField
              control={control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      {...field}
                      type="email"
                      placeholder="yourname@example.com"
                    />
                  </FormControl>
                  <FormDescription>
                    Enter your email address for your account
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input {...field} type="text" placeholder="username" />
                  </FormControl>
                  <FormDescription>
                    Enter your username for your account
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input {...field} type="password" placeholder="password" />
                  </FormControl>
                  <FormDescription>
                    Enter your password (min. 8 characters)
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <div className="text-sm">
              Already have an account?{" "}
              <Button variant="link">
                <Link href="/login">Login here</Link>
              </Button>
            </div>
          </CardContent>

          <CardFooter className="flex flex-col items-center justify-center">
            <Button type="submit" disabled={isPending}>
              Register
            </Button>
          </CardFooter>
        </Card>
      </form>
    </Form>
  );
}
