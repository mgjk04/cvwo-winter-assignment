const refreshURL = 'http://localhost:8080/refresh'

export const refresh = async () => {
    const res = await fetch(refreshURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: 'include'
  });
  if (!res.ok) throw new Error(String(res.status));
}
