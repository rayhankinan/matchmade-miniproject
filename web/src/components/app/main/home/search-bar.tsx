"use client";

import { useCallback, useEffect } from "react";
import { useSearchParams, usePathname, useRouter } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";
import { zodResolver } from "@hookform/resolvers/zod";
import { type SubmitHandler, useForm } from "react-hook-form";
import { z } from "zod";

import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormMessage,
} from "~/components/ui/form";
import { Input } from "~/components/ui/input";

const formSchema = z.object({
  term: z.string(),
});

type SearchFormData = z.infer<typeof formSchema>;

export default function SearchBar() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const router = useRouter();

  const form = useForm<SearchFormData>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      term: searchParams.get("query") ?? "",
    },
    mode: "onBlur",
  });
  const { handleSubmit, watch, control } = form;

  const onSubmit: SubmitHandler<SearchFormData> = useCallback(
    (data) => {
      const term = data.term;
      const params = new URLSearchParams(searchParams);

      if (term) {
        params.set("query", term);
      } else {
        params.delete("query");
      }

      router.replace(`${pathname}?${params.toString()}`);
    },
    [searchParams, pathname, router],
  );

  const onDebouncedChange = useDebouncedCallback(
    () => void handleSubmit(onSubmit)(),
    300,
  );

  useEffect(() => {
    const subscription = watch(onDebouncedChange);
    return () => subscription.unsubscribe();
  }, [watch, onDebouncedChange]);

  return (
    <Form {...form}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <FormField
          control={control}
          name="term"
          render={({ field }) => (
            <FormItem>
              <FormControl>
                <Input
                  {...field}
                  type="text"
                  placeholder="Enter a search term..."
                />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
      </form>
    </Form>
  );
}
