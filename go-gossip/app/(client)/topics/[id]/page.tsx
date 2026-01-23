"use client";
import useReadPost from "@/app/(client)/posts/_hooks/useReadPost";
import {
  Box,
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
import { dataState } from "../types";
import { post } from "@/app/(client)/posts/types";
import useReadTopic from "../_hooks/useReadTopic";
import { getCookie } from "cookies-next";
import { userState } from "@/app/types";
import DeleteIcon from '@mui/icons-material/Delete';
import useDeleteTopic from "../_hooks/useDeleteTopic";
import { useRouter } from "next/navigation";

export default function TopicPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const router = useRouter();
  const [dataState, setDataState] = useState<dataState>({
    page: 1,
    limit: 10,
  });
  const [userState, setUserState] = useState<userState>({
    username: "",
    userId: "",
  });
  useEffect(() => {
    const username = getCookie("username") || "";
    const userId = getCookie("user_id") || "";
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string,
      userId: userId as string,
    });
  }, []);

  const readPostURL =
    process.env.NEXT_PUBLIC_API_URL +
    `/topics/${id}/posts?page=${dataState.page}&limit=${dataState.limit}`;
  const readTopicURL = `${process.env.NEXT_PUBLIC_API_URL}/topics/${id}`;
  const createPostURL = `/topics/${id}/create`;
  const editTopicURL = `/topics/${id}/edit`;
  const deleteTopicURL = readTopicURL;
  const {mutate: deleteTopic} = useDeleteTopic(deleteTopicURL, router);
  const {
    data: topicData,
    isPending: topicIsPending,
    isError: postIsError,
  } = useReadTopic(readTopicURL);
  const { data: postData, isPending, isError } = useReadPost(readPostURL);
  const info = postData || { posts: [], count: 0 };
  const isAuthor = userState?.userId && userState.userId === topicData?.author_id;
  const isSignedIn =  userState?.userId != "";

  return (
    <Stack>
      <Grid container>
        <Grid size="grow">
          <Typography variant="h5">{topicData?.topicname || ""}</Typography>
        </Grid>
        <Grid size="auto">
          {isAuthor && (
            <Button component={Link} href={editTopicURL}>
              Edit Topic
            </Button>
          )}
        </Grid>
        <Grid size="auto">
          {isAuthor && (
            <Button onClick={() => deleteTopic(undefined)}>
              <DeleteIcon/>
            </Button>
          )}
        </Grid>
        <Grid size="auto" justifyContent="center" alignItems="center">
          {isSignedIn && <Button component={Link} href={createPostURL}>
            Make conversation?
          </Button>}
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
