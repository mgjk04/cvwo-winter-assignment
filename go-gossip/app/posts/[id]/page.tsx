"use client";
import { Button, Grid, Link, List, ListItemText, Stack, Typography } from "@mui/material";
import { use, useEffect, useState } from "react";
import useReadPost from "../_hooks/useReadPost";
import useReadComment from "@/app/comments/_hooks/useReadComment";
import { comment } from "@/app/comments/types";
import dayjs from "dayjs";
import { getCookie } from "cookies-next";
import { userState } from "@/app/types";

export default function PostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);
  const [dataState, setDataState] = useState<typeof dataState>({
    page: 1,
    limit: 10,
  });
  const [userState, setUserState] = useState<userState>({
    username: "",
    userId: "",
  });

  useEffect(() => {
    const username = getCookie("username");
    const userId = getCookie("user_id");
    // eslint-disable-next-line react-hooks/set-state-in-effect
    setUserState({
      username: username as string | undefined,
      userId: userId as string | undefined,
    });
  }, []);

  const readPostURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}`;
  const readCommentURL = `${process.env.NEXT_PUBLIC_API_URL}/posts/${id}/comments?page=${dataState.page}&limit=${dataState.limit}`;
  const createCommentURL = `/posts/${id}/create`;
  const editPostURL = `/posts/${id}/edit`;
  const {
    data: postData,
    isPending: postIsPending,
    isError: postIsError,
  } = useReadPost(readPostURL);
  const {
    data: commentData,
    isPending: commentIsPending,
    isError: commentIsError,
  } = useReadComment(readCommentURL);
  const info = commentData || { comments: [], count: 0 };
  return (
    <Stack>
      <Grid container>
        <Grid size="grow">
          <Stack>
            <Typography variant="h5">{postData?.title || ""}</Typography>
            <Typography variant="body1">
              {postData?.description || ""}
            </Typography>
            <Typography variant="body2">{postData?.authorname}</Typography>
          </Stack>
        </Grid>
        <Grid size="auto">
          {userState?.userId && userState.userId === postData?.author_id && (
            <Button component={Link} href={editPostURL}>
              Edit Post
            </Button>
          )}
        </Grid>
        <Grid size="auto" justifyContent="center" alignItems="center">
          <Button component={Link} href={createCommentURL}>
            Comment
          </Button>
        </Grid>
      </Grid>
      <List>
        {info.count === 0 ? (
          <Typography variant="body2">Wow such empty...</Typography>
        ) : (
          info.comments?.map((c: comment) => {
            return (
              <ListItemText key={c.id}>
                <Typography variant="body1">{c.content}</Typography>
                <Typography variant="body2">
                  {dayjs(c.created_at).format("HH:MM DD-MM-YYYY")}
                </Typography>
                <Typography variant="body2">{c.authorname}</Typography>
              </ListItemText>
            );
          })
        )}
      </List>
    </Stack>
  );
}
