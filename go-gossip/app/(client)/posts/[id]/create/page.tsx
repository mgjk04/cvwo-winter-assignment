"use client";
import { use } from "react";
import CreateComment from "./_components/createComment";

export default function Page({params}:{params: Promise<{id: string}>}) {
  const { id } = use(params);
  const createCommentURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}/comments`;
  return <CreateComment submitURL={createCommentURL} />;
}
