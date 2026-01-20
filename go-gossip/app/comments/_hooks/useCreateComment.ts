import { useMutation, useQueryClient } from "@tanstack/react-query";
import z from "zod";
import { commentFormSchema } from "../zod";

const createComment = (createURL: string) => async (values: z.infer<typeof commentFormSchema>) => {
  const res = await fetch(createURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include"
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useCreateComment(createURL: string) {
  const client = useQueryClient();
  return useMutation({
    mutationFn: createComment(createURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["comment"] });
    },
  });
}
