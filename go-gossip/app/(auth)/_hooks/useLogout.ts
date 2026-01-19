import { useMutation } from "@tanstack/react-query";

const logoutURL = "http://localhost:8080/logout";

const logout = async () => {
  const res = await fetch(logoutURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: "include",
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useLogout() {
  return useMutation({
    mutationFn: logout,
  });
}
