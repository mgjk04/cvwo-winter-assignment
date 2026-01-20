"use client";
import useReadPost from "@/app/posts/_hooks/useReadPost";
import {
  Button,
  Grid,
  Link,
  List,
  ListItemButton,
  ListItemText,
  Stack,
  Typography,
} from "@mui/material";
import { use, useEffect, useState } from "react";
import { topic, dataState } from "../types";
import { post } from "@/app/posts/types";
import useReadTopic from "../_hooks/useReadTopic";
import { getCookie } from "cookies-next";
import { userState } from "@/app/types";

export default function TopicPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const [dataState, setDataState] = useState<dataState>({
    page: 1,
    limit: 10,
  });
  const [userState, setUserState] = useState<userState>({username: "", userId: ""});
  useEffect(() => {
    const username = getCookie("username");
    const userId = getCookie("user_id");
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string | undefined,
      userId: userId as string | undefined,
    });
  }, []);

  const readPostURL =
    process.env.NEXT_PUBLIC_API_URL +
    `/topics/${id}/posts?page=${dataState.page}&limit=${dataState.limit}`;
  const readTopicURL = process.env.NEXT_PUBLIC_API_URL + `/topics/${id}`;
  const createPostURL = `/topics/${id}/create`;
  const editTopicURL = `/topics/${id}/edit`;
  const {
    data: topicData,
    isPending: topicIsPending,
    isError: postIsError,
  } = useReadTopic(readTopicURL);
  const { data: postData, isPending, isError } = useReadPost(readPostURL);
  const info = postData || { posts: [], count: 0 };

  
  return (
    <Stack>
      <Grid container>
        <Grid size="grow">
          <Typography variant="h5">{topicData?.topicname || ""}</Typography>
        </Grid>
        <Grid size="auto">
          {userState?.userId && userState.userId=== topicData?.author_id && (
            <Button component={Link} href={editTopicURL}>
              Edit Topic
            </Button>
          )}
        </Grid>
        <Grid size="auto" justifyContent="center" alignItems="center">
          <Button component={Link} href={createPostURL}>
            Make conversation?
          </Button>
        </Grid>
      </Grid>
      <List>
        {info.count === 0 ? (
          <Typography variant="body2">Wow such empty...</Typography>
        ) : (
          info.posts?.map((p: post) => {
            const postURL = `/posts/${p.id}`;
            return (
              <ListItemButton key={p.id} LinkComponent={Link} href={postURL}>
                <ListItemText>
                  <Typography variant="h6">{p.title}</Typography>
                  <Typography variant="body1">{p.description}</Typography>
                  <Typography variant="body2">{p.authorname}</Typography>
                </ListItemText>
              </ListItemButton>
            );
          })
        )}
      </List>
    </Stack>
  );
}
