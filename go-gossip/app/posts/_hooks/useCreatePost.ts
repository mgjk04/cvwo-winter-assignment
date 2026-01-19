import { useMutation, useQueryClient } from "@tanstack/react-query";
import z from "zod";
import { postFormSchema } from "../zod";

const createURL = "http://localhost:8080/posts/";

const createPost = async (values: z.infer<typeof postFormSchema>) => {
  const res = await fetch(createURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include"
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useCreatePost() {
  const client = useQueryClient();
  return useMutation({
    mutationFn: createPost,
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
    },
  });
}
