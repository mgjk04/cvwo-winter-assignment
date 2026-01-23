"use client";
import { use } from "react";
import CreatePost from "./_components/createPost";

export default function Page({params}:{params: Promise<{id: string}>}) {
  const { id } = use(params);
  const createPostURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/${id}/posts`;
  return <CreatePost submitURL={createPostURL} />;
}
