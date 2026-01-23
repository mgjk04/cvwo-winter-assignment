"use client";
import CreateTopic from "./_components/createTopic";

export default function Page() {
  const createTopicURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/`;
  return <CreateTopic submitURL={createTopicURL} />;
}
