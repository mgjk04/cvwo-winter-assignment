import type { Metadata } from "next";
import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";
import "./globals.css";
import Providers from "./providers";
import "dotenv/config";

export const metadata: Metadata = {
  title: "GO-GOSSIP",
  description: "Share your thoughts",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  return (
    <html lang="en">
      <body className={`antialiased`}>
        <Providers>
          {children}
        </Providers>
      </body>
    </html>
  );
}
