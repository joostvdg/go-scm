package webhooks

type OpenPullRequests struct {
    Size       int  `json:"size"`
    Limit      int  `json:"limit"`
    IsLastPage bool `json:"isLastPage"`
    Values     []struct {
        ID          int    `json:"id"`
        Version     int    `json:"version"`
        Title       string `json:"title"`
        Description string `json:"description"`
        State       string `json:"state"`
        Open        bool   `json:"open"`
        Closed      bool   `json:"closed"`
        CreatedDate int64  `json:"createdDate"`
        UpdatedDate int64  `json:"updatedDate"`
        FromRef     struct {
            ID           string `json:"id"`
            DisplayID    string `json:"displayId"`
            LatestCommit string `json:"latestCommit"`
            Repository   struct {
                Slug          string `json:"slug"`
                ID            int    `json:"id"`
                Name          string `json:"name"`
                ScmID         string `json:"scmId"`
                State         string `json:"state"`
                StatusMessage string `json:"statusMessage"`
                Forkable      bool   `json:"forkable"`
                Project       struct {
                    Key    string `json:"key"`
                    ID     int    `json:"id"`
                    Name   string `json:"name"`
                    Public bool   `json:"public"`
                    Type   string `json:"type"`
                    Links  struct {
                        Self []struct {
                            Href string `json:"href"`
                        } `json:"self"`
                    } `json:"links"`
                } `json:"project"`
                Public bool `json:"public"`
                Links  struct {
                    Clone []struct {
                        Href string `json:"href"`
                        Name string `json:"name"`
                    } `json:"clone"`
                    Self []struct {
                        Href string `json:"href"`
                    } `json:"self"`
                } `json:"links"`
            } `json:"repository"`
        } `json:"fromRef"`
        ToRef struct {
            ID           string `json:"id"`
            DisplayID    string `json:"displayId"`
            LatestCommit string `json:"latestCommit"`
            Repository   struct {
                Slug          string `json:"slug"`
                ID            int    `json:"id"`
                Name          string `json:"name"`
                ScmID         string `json:"scmId"`
                State         string `json:"state"`
                StatusMessage string `json:"statusMessage"`
                Forkable      bool   `json:"forkable"`
                Project       struct {
                    Key    string `json:"key"`
                    ID     int    `json:"id"`
                    Name   string `json:"name"`
                    Public bool   `json:"public"`
                    Type   string `json:"type"`
                    Links  struct {
                        Self []struct {
                            Href string `json:"href"`
                        } `json:"self"`
                    } `json:"links"`
                } `json:"project"`
                Public bool `json:"public"`
                Links  struct {
                    Clone []struct {
                        Href string `json:"href"`
                        Name string `json:"name"`
                    } `json:"clone"`
                    Self []struct {
                        Href string `json:"href"`
                    } `json:"self"`
                } `json:"links"`
            } `json:"repository"`
        } `json:"toRef"`
        Locked bool `json:"locked"`
        Author struct {
            User struct {
                Name         string `json:"name"`
                EmailAddress string `json:"emailAddress"`
                ID           int    `json:"id"`
                DisplayName  string `json:"displayName"`
                Active       bool   `json:"active"`
                Slug         string `json:"slug"`
                Type         string `json:"type"`
                Links        struct {
                    Self []struct {
                        Href string `json:"href"`
                    } `json:"self"`
                } `json:"links"`
            } `json:"user"`
            Role     string `json:"role"`
            Approved bool   `json:"approved"`
            Status   string `json:"status"`
        } `json:"author"`
        Reviewers []struct {
            User struct {
                Name         string `json:"name"`
                EmailAddress string `json:"emailAddress"`
                ID           int    `json:"id"`
                DisplayName  string `json:"displayName"`
                Active       bool   `json:"active"`
                Slug         string `json:"slug"`
                Type         string `json:"type"`
                Links        struct {
                    Self []struct {
                        Href string `json:"href"`
                    } `json:"self"`
                } `json:"links"`
            } `json:"user"`
            LastReviewedCommit string `json:"lastReviewedCommit"`
            Role               string `json:"role"`
            Approved           bool   `json:"approved"`
            Status             string `json:"status"`
        } `json:"reviewers"`
        Participants []interface{} `json:"participants"`
        Properties   struct {
            MergeResult struct {
                Outcome string `json:"outcome"`
                Current bool   `json:"current"`
            } `json:"mergeResult"`
            ResolvedTaskCount int `json:"resolvedTaskCount"`
            CommentCount      int `json:"commentCount"`
            OpenTaskCount     int `json:"openTaskCount"`
        } `json:"properties"`
        Links struct {
            Self []struct {
                Href string `json:"href"`
            } `json:"self"`
        } `json:"links"`
    } `json:"values"`
    Start int `json:"start"`
}