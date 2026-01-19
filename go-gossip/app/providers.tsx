"use client";
import { AppRouterCacheProvider } from "@mui/material-nextjs/v16-appRouter";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

export default function Providers({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
    const queryClient = new QueryClient();
  return (
    <AppRouterCacheProvider options={{ enableCssLayer: true }}>
      <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
    </AppRouterCacheProvider>
  );
}
