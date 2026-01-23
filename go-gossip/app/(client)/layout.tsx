"use client"
import AppBar from "./_components/AppBar";
import { Stack } from "@mui/material"

export default function ClientLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <Stack>
      <AppBar />
      {children}
    </Stack>
  );
}
