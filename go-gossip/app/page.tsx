"use client";
import { Grid, Typography, Stack, Card, Link, Button, CardActions } from "@mui/material";

export default function Home() {
  return (
    <Stack>
      <Grid>
        <Card>
          <Typography className="font-black" variant="h1">Go Gossip!</Typography>
          <Card>
            <CardActions>
              <Button variant="contained" component={Link} href='/topics'>Find out about the Topics we share...</Button>
            </CardActions>
          </Card>
        </Card>
      </Grid>
    </Stack>
  );
}
