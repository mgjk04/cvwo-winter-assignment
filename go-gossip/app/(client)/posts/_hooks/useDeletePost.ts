import { useMutation, useQueryClient } from "@tanstack/react-query";
import { AppRouterInstance } from "next/dist/shared/lib/app-router-context.shared-runtime";

const deleteTopic = async (deleteURL: string) => {
  const res = await fetch(deleteURL, {
    method: "DELETE",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useDeletePost(deleteURL: string, router: AppRouterInstance) {
  const client = useQueryClient();
  return useMutation({
    mutationFn: () => deleteTopic(deleteURL),
    onSettled: () => {
      client.invalidateQueries({ queryKey: ["post"] });
      router.back();
    },
  });
}
