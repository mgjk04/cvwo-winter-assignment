import { useMutation, useQueryClient } from "@tanstack/react-query";
import z from "zod";
import { postFormSchema } from "../zod";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

const createPost = (createURL: string) => async (values: z.infer<typeof postFormSchema>) => {
  const res = await fetch(createURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include"
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useCreatePost(createURL: string, router: AppRouterInstance) {
  const client = useQueryClient();
  return useMutation({
    mutationFn: createPost(createURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
      router.back();
    },
  });
}
