const refreshURL = `${process.env.NEXT_PUBLIC_API_URL}/refresh`;

export const refresh = async () => {
    const res = await fetch(refreshURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    credentials: 'include'
  });
  if (!res.ok) throw new Error(String(res.status));
}
