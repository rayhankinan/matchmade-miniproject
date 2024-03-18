"use client";

import { useRouter } from "next/navigation";
import { useMutation } from "@tanstack/react-query";
import { zodResolver } from "@hookform/resolvers/zod";
import { type SubmitHandler, useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "~/components/ui/button";
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
  identifier: z.string(),
  password: z.string(),
});

type LoginFormData = z.infer<typeof formSchema>;

type LoginFormResponse = {
  data: null;
};

export default function LoginForm() {
  const { toast } = useToast();
  const router = useRouter();

  const mutation = useMutation({
    mutationFn: async (data: LoginFormData) =>
      await api.post<LoginFormResponse>("/login", data),
    onSuccess: () => router.push("/"),
    onError: (error) =>
      toast({
        variant: "destructive",
        title: "An error occurred",
        description: error.message,
      }),
  });

  const form = useForm<LoginFormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      identifier: "",
      password: "",
    },
  });

  const onSubmit: SubmitHandler<LoginFormData> = (data) => {
    mutation.mutate(data);
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <FormField
          control={form.control}
          name="identifier"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email or Username</FormLabel>
              <FormControl>
                <Input {...field} type="text" placeholder="username" />
              </FormControl>
              <FormDescription>
                Enter the email address or username associated with your account
              </FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input {...field} type="password" placeholder="password" />
              </FormControl>
              <FormDescription>Enter your password</FormDescription>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">Submit</Button>
      </form>
    </Form>
  );
}
