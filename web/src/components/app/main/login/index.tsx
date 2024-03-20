"use client";

import { useRouter } from "next/navigation";
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
  identifier: z.string(),
  password: z.string(),
});

type LoginFormData = z.infer<typeof formSchema>;

export default function LoginForm() {
  const router = useRouter();

  const { toast } = useToast();

  const form = useForm<LoginFormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      identifier: "",
      password: "",
    },
    mode: "onBlur",
  });

  const { handleSubmit, reset, control } = form;

  const loginMutation = useMutation({
    mutationFn: async (data: LoginFormData) =>
      await api.post("/users/login", data),
    onSuccess: () => {
      router.push("/");
      router.refresh();
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

  const { mutate, isPending } = loginMutation;

  const onSubmit: SubmitHandler<LoginFormData> = (data) => mutate(data);

  return (
    <Form {...form}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Card>
          <CardHeader>
            <CardTitle>Login</CardTitle>
            <CardDescription>Enter your credentials</CardDescription>
          </CardHeader>

          <CardContent className="space-y-4">
            <FormField
              control={control}
              name="identifier"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>Email or Username</FormLabel>
                  <FormControl>
                    <Input {...field} type="text" placeholder="username" />
                  </FormControl>
                  <FormDescription>
                    Enter the email address or username
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
                  <FormDescription>Enter your password</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />
          </CardContent>

          <CardFooter>
            <Button type="submit" disabled={isPending}>
              Submit
            </Button>
          </CardFooter>
        </Card>
      </form>
    </Form>
  );
}
