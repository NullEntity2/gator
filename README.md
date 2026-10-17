# gator

Gator is a cli RSS blog aggregator. 

## Installation

Gator requires Go and a Postgres database.
Create ~/.gatorconfig with the following contents:
```json
{
    "db_url":"postgres://bryce:@localhost:5432/gator?sslmode=disable"
}
```
Update the db_url to point to your own database.

The goose dependency is required:
`go install github.com/pressly/goose/v3/cmd/goose@latest`

cd into `sql/schema` and run `goose postgres <connection string> up`

## Usage
1. `gator reset`
2. `gator register bryce`
3. `gator addfeed addfeed "Hacker News RSS" "https://hnrss.org/newest"`
4. `gator run agg 30s`
5. Let it run for a bit to scan feeds, then ctrl+c
6. `gator run browse`


## Command Reference
* addfeed *NAME* *URL*
* agg *TIMESPAN* - Timespan is expected to be an interval like 30s, 5m, 1h. Loads posts from feeds from all users.
* browse [PAGE_SIZE] - Displays the latest *PAGE_SIZE* articles. Defaults to 2.
* feeds - Lists all available feeds
* follow *URL* - Follows a feed for the current user
* following - List all feeds the current user is following
* login *NAME* - Logs in as the provided user
* register *NAME* - Creates an account for the given user and switches the logged in user.
* reset - Deletes everything from the database
* unfollow *URL* - Unfollows a feed for the current user
* users - List all registered users