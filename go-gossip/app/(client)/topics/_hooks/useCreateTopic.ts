import { useMutation, useQueryClient } from "@tanstack/react-query";
import z from "zod";
import { topicFormSchema } from "../zod";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

const createURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/`;

const createTopic = async (values: z.infer<typeof topicFormSchema>) => {
  const res = await fetch(createURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include"
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useCreateTopic( router: AppRouterInstance) {
  const client = useQueryClient();
  return useMutation({
    mutationFn: createTopic,
    onSettled: () => {
      client.invalidateQueries({queryKey:['topic']})
      router.back();
    }
  });
}
