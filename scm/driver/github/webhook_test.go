// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/joostvdg/go-scm/scm"
	"github.com/stretchr/testify/assert"

	"github.com/google/go-cmp/cmp"
)

func TestWebhooks(t *testing.T) {
	tests := []struct {
		event  string
		before string
		after  string
		obj    interface{}
	}{
		//
		// push events
		//

		// fork
		{
			event:  "fork",
			before: "testdata/webhooks/fork.json",
			after:  "testdata/webhooks/fork.json.golden",
			obj:    new(scm.ForkHook),
		},

		// repository
		{
			event:  "repository",
			before: "testdata/webhooks/repository.json",
			after:  "testdata/webhooks/repository.json.golden",
			obj:    new(scm.RepositoryHook),
		},

		// installation_repositories
		{
			event:  "installation_repositories",
			before: "testdata/webhooks/installation_repository.json",
			after:  "testdata/webhooks/installation_repository.json.golden",
			obj:    new(scm.InstallationRepositoryHook),
		},

		// check_suite
		{
			event:  "check_suite",
			before: "testdata/webhooks/check_suite_created.json",
			after:  "testdata/webhooks/check_suite_created.json.golden",
			obj:    new(scm.CheckSuiteHook),
		},

		// deployment_status
		{
			event:  "deployment_status",
			before: "testdata/webhooks/deployment_status.json",
			after:  "testdata/webhooks/deployment_status.json.golden",
			obj:    new(scm.DeploymentStatusHook),
		},

		// release
		{
			event:  "release",
			before: "testdata/webhooks/release.json",
			after:  "testdata/webhooks/release.json.golden",
			obj:    new(scm.ReleaseHook),
		},

		// status
		{
			event:  "status",
			before: "testdata/webhooks/status.json",
			after:  "testdata/webhooks/status.json.golden",
			obj:    new(scm.StatusHook),
		},

		// label
		{
			event:  "label",
			before: "testdata/webhooks/label_deleted.json",
			after:  "testdata/webhooks/label_deleted.json.golden",
			obj:    new(scm.LabelHook),
		},

		// ping
		{
			event:  "ping",
			before: "testdata/webhooks/ping.json",
			after:  "testdata/webhooks/ping.json.golden",
			obj:    new(scm.PingHook),
		},

		// push hooks
		{
			event:  "push",
			before: "testdata/webhooks/push.json",
			after:  "testdata/webhooks/push.json.golden",
			obj:    new(scm.PushHook),
		},
		// push tag create hooks
		{
			event:  "push",
			before: "testdata/webhooks/push_tag.json",
			after:  "testdata/webhooks/push_tag.json.golden",
			obj:    new(scm.PushHook),
		},
		// push tag delete hooks
		{
			event:  "push",
			before: "testdata/webhooks/push_tag_delete.json",
			after:  "testdata/webhooks/push_tag_delete.json.golden",
			obj:    new(scm.PushHook),
		},
		// push branch create
		{
			event:  "push",
			before: "testdata/webhooks/push_branch_create.json",
			after:  "testdata/webhooks/push_branch_create.json.golden",
			obj:    new(scm.PushHook),
		},
		// push branch delete
		{
			event:  "push",
			before: "testdata/webhooks/push_branch_delete.json",
			after:  "testdata/webhooks/push_branch_delete.json.golden",
			obj:    new(scm.PushHook),
		},

		//
		// branch events
		//

		// push branch create
		{
			event:  "create",
			before: "testdata/webhooks/branch_create.json",
			after:  "testdata/webhooks/branch_create.json.golden",
			obj:    new(scm.BranchHook),
		},
		// push branch delete
		{
			event:  "delete",
			before: "testdata/webhooks/branch_delete.json",
			after:  "testdata/webhooks/branch_delete.json.golden",
			obj:    new(scm.BranchHook),
		},

		//
		// tag events
		//

		// push tag create
		{
			event:  "create",
			before: "testdata/webhooks/tag_create.json",
			after:  "testdata/webhooks/tag_create.json.golden",
			obj:    new(scm.TagHook),
		},
		// push tag delete
		{
			event:  "delete",
			before: "testdata/webhooks/tag_delete.json",
			after:  "testdata/webhooks/tag_delete.json.golden",
			obj:    new(scm.TagHook),
		},

		//
		// pull request events
		//

		// pull request synced
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_sync.json",
			after:  "testdata/webhooks/pr_sync.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request opened
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_opened.json",
			after:  "testdata/webhooks/pr_opened.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request closed
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_closed.json",
			after:  "testdata/webhooks/pr_closed.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request reopened
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_reopened.json",
			after:  "testdata/webhooks/pr_reopened.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request edited
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_edited.json",
			after:  "testdata/webhooks/pr_edited.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request labeled
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_labeled.json",
			after:  "testdata/webhooks/pr_labeled.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request unlabeled
		{
			event:  "pull_request",
			before: "testdata/webhooks/pr_unlabeled.json",
			after:  "testdata/webhooks/pr_unlabeled.json.golden",
			obj:    new(scm.PullRequestHook),
		},
		// pull request comment
		{
			event:  "pull_request_review_comment",
			before: "testdata/webhooks/pr_comment.json",
			after:  "testdata/webhooks/pr_comment.json.golden",
			obj:    new(scm.PullRequestCommentHook),
		},
		// issue comment
		{
			event:  "issue_comment",
			before: "testdata/webhooks/issue_comment.json",
			after:  "testdata/webhooks/issue_comment.json.golden",
			obj:    new(scm.IssueCommentHook),
		},
		// deployment
		{
			event:  "deployment",
			before: "testdata/webhooks/deployment.json",
			after:  "testdata/webhooks/deployment.json.golden",
			obj:    new(scm.DeployHook),
		},

		// installation of GitHub App
		{
			event:  "installation",
			before: "testdata/webhooks/installation.json",
			after:  "testdata/webhooks/installation.json.golden",
			obj:    new(scm.InstallationHook),
		},
		// delete installation of GitHub App
		{
			event:  "installation",
			before: "testdata/webhooks/installation_delete.json",
			after:  "testdata/webhooks/installation_delete.json.golden",
			obj:    new(scm.InstallationHook),
		},
	}

	for _, test := range tests {
		before, err := ioutil.ReadFile(test.before)
		if err != nil {
			t.Error(err)
			continue
		}
		after, err := ioutil.ReadFile(test.after)
		if err != nil {
			t.Error(err)
			continue
		}

		buf := bytes.NewBuffer(before)
		r, _ := http.NewRequest("GET", "/", buf)
		r.Header.Set("X-GitHub-Event", test.event)
		r.Header.Set("X-Hub-Signature", "sha1=380f462cd2e160b84765144beabdad2e930a7ec5")
		r.Header.Set("X-GitHub-Delivery", "f2467dea-70d6-11e8-8955-3c83993e0aef")

		s := new(webhookService)
		o, err := s.Parse(r, secretFunc)
		if err != nil && err != scm.ErrSignatureInvalid {
			t.Logf("failed to parse webhook for test %s", test.event)
			t.Error(err)
			continue
		}

		err = json.Unmarshal(after, test.obj)
		if err != nil {
			t.Error(err, "failed to unmarshal", test.after, "for test", test.event)
			continue
		}

		if diff := cmp.Diff(test.obj, o); diff != "" {
			t.Errorf("Error unmarshaling %s", test.before)
			t.Log(diff)

			// debug only. remove once implemented
			json.NewEncoder(os.Stdout).Encode(o)
		}

		switch event := o.(type) {
		case *scm.PushHook:
			if !strings.HasPrefix(event.Ref, "refs/") {
				t.Errorf("Push hook reference must start with refs/")
			}
		case *scm.BranchHook:
			if strings.HasPrefix(event.Ref.Name, "refs/") {
				t.Errorf("Branch hook reference must not start with refs/")
			}
		case *scm.TagHook:
			if strings.HasPrefix(event.Ref.Name, "refs/") {
				t.Errorf("Branch hook reference must not start with refs/")
			}
		case *scm.InstallationHook:
			assert.NotNil(t, event.Installation, "InstallationHook.Installation")
			assert.NotNil(t, event.GetInstallationRef(), "InstallationHook.GetInstallationRef()")
			assert.NotEmpty(t, event.GetInstallationRef().ID, "InstallationHook.GetInstallationRef().ID")
		}
	}
}

func TestWebhook_ErrUnknownEvent(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/webhooks/push.json")
	r, _ := http.NewRequest("GET", "/", bytes.NewBuffer(f))
	r.Header.Set("X-GitHub-Delivery", "ee8d97b4-1479-43f1-9cac-fbbd1b80da55")
	r.Header.Set("X-Hub-Signature", "sha1=380f462cd2e160b84765144beabdad2e930a7ec5")

	s := new(webhookService)
	_, err := s.Parse(r, secretFunc)
	if !scm.IsUnknownWebhook(err) {
		t.Errorf("Expect unknown event error, got %v", err)
	}
}

func TestWebhookInvalid(t *testing.T) {
	f, _ := ioutil.ReadFile("testdata/webhooks/push.json")
	r, _ := http.NewRequest("GET", "/", bytes.NewBuffer(f))
	r.Header.Set("X-GitHub-Event", "push")
	r.Header.Set("X-GitHub-Delivery", "ee8d97b4-1479-43f1-9cac-fbbd1b80da55")
	r.Header.Set("X-Hub-Signature", "sha1=380f462cd2e160b84765144beabdad2e930a7ec5")

	s := new(webhookService)
	_, err := s.Parse(r, secretFunc)
	if err != scm.ErrSignatureInvalid {
		t.Errorf("Expect invalid signature error, got %v", err)
	}
}

func TestWebhookValid(t *testing.T) {
	// the sha can be recalculated with the below command
	// openssl dgst -sha1 -hmac <secret> <file>

	f, _ := ioutil.ReadFile("testdata/webhooks/push.json")
	r, _ := http.NewRequest("GET", "/", bytes.NewBuffer(f))
	r.Header.Set("X-GitHub-Event", "push")
	r.Header.Set("X-GitHub-Delivery", "ee8d97b4-1479-43f1-9cac-fbbd1b80da55")
	r.Header.Set("X-Hub-Signature", "sha1=e9c4409d39729236fda483f22e7fb7513e5cd273")

	s := new(webhookService)
	_, err := s.Parse(r, secretFunc)
	if err != nil {
		t.Errorf("Expect valid signature, got %v", err)
	}
}

func secretFunc(scm.Webhook) (string, error) {
	return "topsecret", nil
}
