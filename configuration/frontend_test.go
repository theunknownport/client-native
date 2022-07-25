// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package configuration

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/theunknownport/client-native/v4/misc"
	"github.com/theunknownport/client-native/v4/models"
)

func TestGetFrontends(t *testing.T) { //nolint:gocognit
	v, frontends, err := clientTest.GetFrontends("")
	if err != nil {
		t.Error(err.Error())
	}

	if len(frontends) != 2 {
		t.Errorf("%v frontends returned, expected 2", len(frontends))
	}

	if v != version {
		t.Errorf("Version %v returned, expected %v", v, version)
	}

	for _, f := range frontends {
		if f.Name != "test" && f.Name != "test_2" {
			t.Errorf("Expected only test or test_2 frontend, %v found", f.Name)
		}
		optionValue := "disabled"
		if f.Name == "test" {
			optionValue = "enabled"
			if f.BindProcess != "odd" {
				t.Errorf("%v: BindProcess not all: %v", f.Name, f.BindProcess)
			}
		}
		if f.Mode != "http" {
			t.Errorf("%v: Mode not http: %v", f.Name, f.Mode)
		}
		if f.Dontlognull != "enabled" {
			t.Errorf("%v: Dontlognull not enabled: %v", f.Name, f.Dontlognull)
		}
		if f.HTTPConnectionMode != "httpclose" {
			t.Errorf("%v: HTTPConnectionMode not httpclose: %v", f.Name, f.HTTPConnectionMode)
		}
		if f.Contstats != "enabled" {
			t.Errorf("%v: Contstats not enabled: %v", f.Name, f.Contstats)
		}
		if *f.HTTPRequestTimeout != 2000 {
			t.Errorf("%v: HTTPRequestTimeout not 2: %v", f.Name, *f.HTTPRequestTimeout)
		}
		if *f.HTTPKeepAliveTimeout != 3000 {
			t.Errorf("%v: HTTPKeepAliveTimeout not 3: %v", f.Name, *f.HTTPKeepAliveTimeout)
		}
		if f.DefaultBackend != "test" && f.DefaultBackend != "test_2" {
			t.Errorf("%v: DefaultFarm not test or test_2: %v", f.Name, f.DefaultBackend)
		}
		if *f.Maxconn != 2000 {
			t.Errorf("%v: Maxconn not 2000: %v", f.Name, *f.Maxconn)
		}
		if *f.Backlog != 2048 {
			t.Errorf("%v: Backlog not 2048: %v", f.Name, *f.Backlog)
		}
		if *f.ClientTimeout != 4000 {
			t.Errorf("%v: ClientTimeout not 4: %v", f.Name, *f.ClientTimeout)
		}
		if f.Tcpka != "enabled" {
			t.Errorf("%v: Tcpka not enabled: %v", f.Name, f.Tcpka)
		}
		if f.Clitcpka != "enabled" {
			t.Errorf("%v: Clitcpka not enabled: %v", f.Name, f.Clitcpka)
		}
		if f.ClitcpkaCnt == nil {
			t.Errorf("%v: ClitcpkaCnt is nil", f.Name)
		} else if *f.ClitcpkaCnt != 10 {
			t.Errorf("%v: ClitcpkaCnt not 10: %v", f.Name, *f.ClitcpkaCnt)
		}
		if f.ClitcpkaIdle == nil {
			t.Errorf("%v: ClitcpkaIdle is nil", f.Name)
		} else if *f.ClitcpkaIdle != 10000 {
			t.Errorf("%v: ClitcpkaIdle not 10000: %v", f.Name, *f.ClitcpkaIdle)
		}
		if f.ClitcpkaIntvl == nil {
			t.Errorf("%v: ClitcpkaIntvl is nil", f.Name)
		} else if *f.ClitcpkaIntvl != 10000 {
			t.Errorf("%v: ClitcpkaIntvl not 10000: %v", f.Name, *f.ClitcpkaIntvl)
		}
		if f.HTTPIgnoreProbes != optionValue {
			t.Errorf("%v: HTTPIgnoreProbes not %s: %v", f.Name, optionValue, f.HTTPIgnoreProbes)
		}
		if f.HTTPUseProxyHeader != optionValue {
			t.Errorf("%v: HTTPUseProxyHeader not %s: %v", f.Name, optionValue, f.HTTPUseProxyHeader)
		}
		if f.Httpslog != optionValue {
			t.Errorf("%v: Httpslog not %s: %v", f.Name, optionValue, f.Httpslog)
		}
		if f.IndependentStreams != optionValue {
			t.Errorf("%v: IndependentStreams not %s: %v", f.Name, optionValue, f.IndependentStreams)
		}
		if f.Nolinger != optionValue {
			t.Errorf("%v: Nolinger not %s: %v", f.Name, optionValue, f.Nolinger)
		}
		if f.Originalto != optionValue {
			t.Errorf("%v: Originalto not %s: %v", f.Name, optionValue, f.Originalto)
		}
		if f.SocketStats != optionValue {
			t.Errorf("%v: SocketStats not %s: %v", f.Name, optionValue, f.SocketStats)
		}
		if f.TCPSmartAccept != optionValue {
			t.Errorf("%v: TCPSmartAccept not %s: %v", f.Name, optionValue, f.TCPSmartAccept)
		}
		if f.DontlogNormal != optionValue {
			t.Errorf("%v: DontlogNormal not %s: %v", f.Name, optionValue, f.DontlogNormal)
		}
		if f.HTTPNoDelay != optionValue {
			t.Errorf("%v: HTTPNoDelay not %s: %v", f.Name, optionValue, f.HTTPNoDelay)
		}
		if f.SpliceAuto != optionValue {
			t.Errorf("%v: SpliceAuto not %s: %v", f.Name, optionValue, f.SpliceAuto)
		}
		if f.SpliceRequest != optionValue {
			t.Errorf("%v: SpliceRequest not %s: %v", f.Name, optionValue, f.SpliceRequest)
		}
		if f.SpliceResponse != optionValue {
			t.Errorf("%v: SpliceResponse not %s: %v", f.Name, optionValue, f.SpliceResponse)
		}
		if f.IdleCloseOnResponse != optionValue {
			t.Errorf("%v: IdleCloseOnResponse not %s: %v", f.Name, optionValue, f.IdleCloseOnResponse)
		}
		if f.StatsOptions == nil {
			t.Errorf("%v: StatsOptions is nil", f.Name)
		}
		if f.StatsOptions.StatsShowModules != true {
			t.Error("StatsShowModules not set")
		}
		if f.StatsOptions.StatsRealm != true {
			t.Error("StatsRealm not set")
		}
		if f.StatsOptions.StatsRealmRealm == nil {
			t.Errorf("%v: StatsRealmRealm is nil", f.Name)
		} else if *f.StatsOptions.StatsRealmRealm != `HAProxy\\ Statistics` {
			t.Errorf("%v: StatsRealmRealm not 'HAProxy Statistics': %v", f.Name, *f.StatsOptions.StatsRealmRealm)
		}
		if len(f.StatsOptions.StatsAuths) != 2 {
			t.Errorf("%v: StatsAuths expected 2 instances got: %v", f.Name, len(f.StatsOptions.StatsAuths))
		}
		if f.StatsOptions.StatsAuths[0].User == nil {
			t.Errorf("%v: StatsAuths 0 User is nil", f.Name)
		} else if *f.StatsOptions.StatsAuths[0].User != "admin" {
			t.Errorf("%v: StatsAuths 0 User not admin: %v", f.Name, *f.StatsOptions.StatsAuths[0].User)
		}
		if f.StatsOptions.StatsAuths[0].Passwd == nil {
			t.Errorf("%v: StatsAuths 0 Passwd is nil", f.Name)
		} else if *f.StatsOptions.StatsAuths[0].Passwd != "AdMiN123" {
			t.Errorf("%v: StatsAuths 0 Passwd not AdMiN123: %v", f.Name, *f.StatsOptions.StatsAuths[0].Passwd)
		}
		if f.StatsOptions.StatsAuths[1].User == nil {
			t.Errorf("%v: StatsAuths 1 User is nil", f.Name)
		} else if *f.StatsOptions.StatsAuths[1].User != "admin2" {
			t.Errorf("%v: StatsAuths 1 User not admin2: %v", f.Name, *f.StatsOptions.StatsAuths[1].User)
		}
		if f.StatsOptions.StatsAuths[1].Passwd == nil {
			t.Errorf("%v: StatsAuths 1 Passwd is nil", f.Name)
		} else if *f.StatsOptions.StatsAuths[1].Passwd != "AdMiN1234" {
			t.Errorf("%v: StatsAuths 1 Passwd not AdMiN1234: %v", f.Name, *f.StatsOptions.StatsAuths[1].Passwd)
		}

	}
}

func TestGetFrontend(t *testing.T) {
	v, f, err := clientTest.GetFrontend("test", "")
	if err != nil {
		t.Error(err.Error())
	}

	if v != version {
		t.Errorf("Version %v returned, expected %v", v, version)
	}

	if f.Name != "test" {
		t.Errorf("Expected only test, %v found", f.Name)
	}
	if f.BindProcess != "odd" {
		t.Errorf("%v: BindProcess not all: %v", f.Name, f.BindProcess)
	}
	if f.Mode != "http" {
		t.Errorf("%v: Mode not http: %v", f.Name, f.Mode)
	}
	if f.MonitorURI != "/healthz" {
		t.Errorf("%v: MonitorURI not /healthz: %v", f.Name, f.MonitorURI)
	}
	if f.MonitorFail == nil {
		t.Errorf("%v: MonitorFail is nil", f.Name)
	}
	if *f.MonitorFail.Cond != "if" {
		t.Errorf("%v: MonitorFail condition not if: %v", f.Name, *f.MonitorFail.Cond)
	}
	if *f.MonitorFail.CondTest != "site_dead" {
		t.Errorf("%v: MonitorFail condition test not site_dead: %v", f.Name, *f.MonitorFail.CondTest)
	}
	if f.Dontlognull != "enabled" {
		t.Errorf("%v: Dontlognull not enabled: %v", f.Name, f.Dontlognull)
	}
	if f.HTTPConnectionMode != "httpclose" {
		t.Errorf("%v: HTTPConnectionMode not httpclose: %v", f.Name, f.HTTPConnectionMode)
	}
	if f.Contstats != "enabled" {
		t.Errorf("%v: Contstats not enabled: %v", f.Name, f.Contstats)
	}
	if *f.HTTPRequestTimeout != 2000 {
		t.Errorf("%v: HTTPRequestTimeout not 2000: %v", f.Name, *f.HTTPRequestTimeout)
	}
	if *f.HTTPKeepAliveTimeout != 3000 {
		t.Errorf("%v: HTTPKeepAliveTimeout not 3000: %v", f.Name, *f.HTTPKeepAliveTimeout)
	}
	if f.DefaultBackend != "test" {
		t.Errorf("%v: DefaultBackend not test: %v", f.Name, f.DefaultBackend)
	}
	if *f.Maxconn != 2000 {
		t.Errorf("%v: Maxconn not 2000: %v", f.Name, *f.Maxconn)
	}
	if *f.Backlog != 2048 {
		t.Errorf("%v: Backlog not 2048: %v", f.Name, *f.Backlog)
	}
	if *f.ClientTimeout != 4000 {
		t.Errorf("%v: ClientTimeout not 4000: %v", f.Name, *f.ClientTimeout)
	}
	if f.AcceptInvalidHTTPRequest != "disabled" {
		t.Errorf("%v: AcceptInvalidHTTPRequest not disabled: %v", f.Name, f.AcceptInvalidHTTPRequest)
	}
	if f.H1CaseAdjustBogusClient != "disabled" {
		t.Errorf("%v: H1CaseAdjustBogusClient not disabled: %v", f.Name, f.H1CaseAdjustBogusClient)
	}
	if f.Tcpka != "enabled" {
		t.Errorf("%v: Tcpka not enabled: %v", f.Name, f.Tcpka)
	}
	if f.Clitcpka != "enabled" {
		t.Errorf("%v: Clitcpka not enabled: %v", f.Name, f.Clitcpka)
	}
	if f.ClitcpkaCnt == nil {
		t.Errorf("%v: ClitcpkaCnt is nil", f.Name)
	} else if *f.ClitcpkaCnt != 10 {
		t.Errorf("%v: ClitcpkaCnt not 10: %v", f.Name, *f.ClitcpkaCnt)
	}
	if f.ClitcpkaIdle == nil {
		t.Errorf("%v: ClitcpkaIdle is nil", f.Name)
	} else if *f.ClitcpkaIdle != 10000 {
		t.Errorf("%v: ClitcpkaIdle not 10000: %v", f.Name, *f.ClitcpkaIdle)
	}
	if f.ClitcpkaIntvl == nil {
		t.Errorf("%v: ClitcpkaIntvl is nil", f.Name)
	} else if *f.ClitcpkaIntvl != 10000 {
		t.Errorf("%v: ClitcpkaIntvl not 10000: %v", f.Name, *f.ClitcpkaIntvl)
	}
	if f.Compression == nil {
		t.Errorf("%v: Compression is nil", f.Name)
	} else {
		if len(f.Compression.Algorithms) != 2 {
			t.Errorf("%v: len Compression.Algorithms not 2: %v", f.Name, len(f.Compression.Algorithms))
		} else {
			if !(f.Compression.Algorithms[0] == "identity" || f.Compression.Algorithms[0] != "gzip") {
				t.Errorf("%v: Compression.Algorithms[0] wrong: %v", f.Name, f.Compression.Algorithms[0])
			}
			if !(f.Compression.Algorithms[1] != "identity" || f.Compression.Algorithms[0] != "gzip") {
				t.Errorf("%v: Compression.Algorithms[1] wrong: %v", f.Name, f.Compression.Algorithms[1])
			}
		}
		if len(f.Compression.Types) != 1 {
			t.Errorf("%v: len Compression.Types not 1: %v", f.Name, len(f.Compression.Types))
		} else {
			if f.Compression.Types[0] != "text/plain" {
				t.Errorf("%v: Compression.Types[0] wrong: %v", f.Name, f.Compression.Types[0])
			}
		}
		if !f.Compression.Offload {
			t.Errorf("%v: Compression.Offload wrong: %v", f.Name, f.Compression.Offload)
		}
	}

	if f.HTTPIgnoreProbes != "enabled" {
		t.Errorf("%v: HTTPIgnoreProbes not enablesd: %v", f.Name, f.HTTPIgnoreProbes)
	}
	if f.HTTPUseProxyHeader != "enabled" {
		t.Errorf("%v: HTTPUseProxyHeader not enablesd: %v", f.Name, f.HTTPUseProxyHeader)
	}
	if f.Httpslog != "enabled" {
		t.Errorf("%v: Httpslog not enablesd: %v", f.Name, f.Httpslog)
	}
	if f.IndependentStreams != "enabled" {
		t.Errorf("%v: IndependentStreams not enablesd: %v", f.Name, f.IndependentStreams)
	}
	if f.Nolinger != "enabled" {
		t.Errorf("%v: Nolinger not enablesd: %v", f.Name, f.Nolinger)
	}
	if f.Originalto != "enabled" {
		t.Errorf("%v: Originalto not enablesd: %v", f.Name, f.Originalto)
	}
	if f.SocketStats != "enabled" {
		t.Errorf("%v: SocketStats not enablesd: %v", f.Name, f.SocketStats)
	}
	if f.TCPSmartAccept != "enabled" {
		t.Errorf("%v: TCPSmartAccept not enablesd: %v", f.Name, f.TCPSmartAccept)
	}
	if f.DontlogNormal != "enabled" {
		t.Errorf("%v: DontlogNormal not enablesd: %v", f.Name, f.DontlogNormal)
	}
	if f.HTTPNoDelay != "enabled" {
		t.Errorf("%v: HTTPNoDelay not enablesd: %v", f.Name, f.HTTPNoDelay)
	}
	if f.SpliceAuto != "enabled" {
		t.Errorf("%v: SpliceAuto not enablesd: %v", f.Name, f.SpliceAuto)
	}
	if f.SpliceRequest != "enabled" {
		t.Errorf("%v: SpliceRequest not enablesd: %v", f.Name, f.SpliceRequest)
	}
	if f.SpliceResponse != "enabled" {
		t.Errorf("%v: SpliceResponse not enablesd: %v", f.Name, f.SpliceResponse)
	}
	if f.IdleCloseOnResponse != "enabled" {
		t.Errorf("%v: IdleCloseOnResponse not enablesd: %v", f.Name, f.IdleCloseOnResponse)
	}
	if f.StatsOptions == nil {
		t.Errorf("%v: StatsOptions is nil", f.Name)
	}
	if f.StatsOptions.StatsShowModules != true {
		t.Error("StatsShowModules not set")
	}
	if f.StatsOptions.StatsRealm != true {
		t.Error("StatsRealm not set")
	}
	if f.StatsOptions.StatsRealmRealm == nil {
		t.Errorf("%v: StatsRealmRealm is nil", f.Name)
	} else if *f.StatsOptions.StatsRealmRealm != `HAProxy\\ Statistics` {
		t.Errorf("%v: StatsRealmRealm not 'HAProxy Statistics': %v", f.Name, *f.StatsOptions.StatsRealmRealm)
	}
	if len(f.StatsOptions.StatsAuths) != 2 {
		t.Errorf("%v: StatsAuths expected 2 instances got: %v", f.Name, len(f.StatsOptions.StatsAuths))
	}
	if f.StatsOptions.StatsAuths[0].User == nil {
		t.Errorf("%v: StatsAuths 0 User is nil", f.Name)
	} else if *f.StatsOptions.StatsAuths[0].User != "admin" {
		t.Errorf("%v: StatsAuths 0 User not admin: %v", f.Name, *f.StatsOptions.StatsAuths[0].User)
	}
	if f.StatsOptions.StatsAuths[0].Passwd == nil {
		t.Errorf("%v: StatsAuths 0 Passwd is nil", f.Name)
	} else if *f.StatsOptions.StatsAuths[0].Passwd != "AdMiN123" {
		t.Errorf("%v: StatsAuths 0 Passwd not AdMiN123: %v", f.Name, *f.StatsOptions.StatsAuths[0].Passwd)
	}
	if f.StatsOptions.StatsAuths[1].User == nil {
		t.Errorf("%v: StatsAuths 1 User is nil", f.Name)
	} else if *f.StatsOptions.StatsAuths[1].User != "admin2" {
		t.Errorf("%v: StatsAuths 1 User not admin2: %v", f.Name, *f.StatsOptions.StatsAuths[1].User)
	}
	if f.StatsOptions.StatsAuths[1].Passwd == nil {
		t.Errorf("%v: StatsAuths 1 Passwd is nil", f.Name)
	} else if *f.StatsOptions.StatsAuths[1].Passwd != "AdMiN1234" {
		t.Errorf("%v: StatsAuths 1 Passwd not AdMiN1234: %v", f.Name, *f.StatsOptions.StatsAuths[1].Passwd)
	}

	_, err = f.MarshalBinary()
	if err != nil {
		t.Error(err.Error())
	}

	_, _, err = clientTest.GetFrontend("doesnotexist", "")
	if err == nil {
		t.Error("Should throw error, non existant frontend")
	}
}

func TestCreateEditDeleteFrontend(t *testing.T) {
	// TestCreateFrontend
	mConn := int64(3000)
	tOut := int64(2)
	clitcpkaCnt := int64(10)
	clitcpkaTimeout := int64(10000)
	statsRealm := "Haproxy Stats"
	f := &models.Frontend{
		Name:                     "created",
		Mode:                     "tcp",
		Maxconn:                  &mConn,
		Httplog:                  true,
		HTTPConnectionMode:       "http-keep-alive",
		HTTPKeepAliveTimeout:     &tOut,
		BindProcess:              "4",
		Logasap:                  "disabled",
		UniqueIDFormat:           "%{+X}o_%fi:%fp_%Ts_%rt:%pid",
		UniqueIDHeader:           "X-Unique-Id",
		AcceptInvalidHTTPRequest: "enabled",
		DisableH2Upgrade:         "enabled",
		ClitcpkaCnt:              &clitcpkaCnt,
		ClitcpkaIdle:             &clitcpkaTimeout,
		ClitcpkaIntvl:            &clitcpkaTimeout,
		HTTPIgnoreProbes:         "enabled",
		HTTPUseProxyHeader:       "enabled",
		Httpslog:                 "enabled",
		IndependentStreams:       "enabled",
		Nolinger:                 "enabled",
		Originalto:               "enabled",
		SocketStats:              "enabled",
		TCPSmartAccept:           "enabled",
		DontlogNormal:            "enabled",
		HTTPNoDelay:              "enabled",
		SpliceAuto:               "enabled",
		SpliceRequest:            "enabled",
		SpliceResponse:           "enabled",
		IdleCloseOnResponse:      "enabled",
		StatsOptions: &models.StatsOptions{
			StatsShowModules: true,
			StatsRealm:       true,
			StatsRealmRealm:  &statsRealm,
			StatsAuths: []*models.StatsAuth{
				{User: misc.StringP("user1"), Passwd: misc.StringP("pwd1")},
				{User: misc.StringP("user2"), Passwd: misc.StringP("pwd2")},
			},
		},
	}

	err := clientTest.CreateFrontend(f, "", version)
	if err != nil {
		t.Error(err.Error())
	} else {
		version++
	}

	v, frontend, err := clientTest.GetFrontend("created", "")
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(frontend, f) {
		fmt.Printf("Created frontend: %v\n", frontend)
		fmt.Printf("Given frontend: %v\n", f)
		t.Error("Created frontend not equal to given frontend")
	}

	if v != version {
		t.Errorf("Version %v returned, expected %v", v, version)
	}

	err = clientTest.CreateFrontend(f, "", version)
	if err == nil {
		t.Error("Should throw error frontend already exists")
		version++
	}

	// TestEditFrontend
	mConn = int64(4000)
	f = &models.Frontend{
		Name:               "created",
		Mode:               "tcp",
		Maxconn:            &mConn,
		Backlog:            misc.Int64P(1024),
		Clflog:             true,
		HTTPConnectionMode: "httpclose",
		BindProcess:        "3",
		MonitorURI:         "/healthz",
		MonitorFail: &models.MonitorFail{
			Cond:     misc.StringP("if"),
			CondTest: misc.StringP("site_is_dead"),
		},
		Compression: &models.Compression{
			Offload: true,
		},
		ClitcpkaCnt:         &clitcpkaCnt,
		ClitcpkaIdle:        &clitcpkaTimeout,
		ClitcpkaIntvl:       &clitcpkaTimeout,
		HTTPIgnoreProbes:    "disabled",
		HTTPUseProxyHeader:  "disabled",
		Httpslog:            "disabled",
		IndependentStreams:  "disabled",
		Nolinger:            "disabled",
		Originalto:          "disabled",
		SocketStats:         "disabled",
		TCPSmartAccept:      "disabled",
		DontlogNormal:       "disabled",
		HTTPNoDelay:         "disabled",
		SpliceAuto:          "disabled",
		SpliceRequest:       "disabled",
		SpliceResponse:      "disabled",
		IdleCloseOnResponse: "disabled",
		StatsOptions: &models.StatsOptions{
			StatsShowModules: true,
			StatsRealm:       true,
			StatsRealmRealm:  &statsRealm,
			StatsAuths: []*models.StatsAuth{
				{User: misc.StringP("new_user1"), Passwd: misc.StringP("new_pwd1")},
				{User: misc.StringP("new_user2"), Passwd: misc.StringP("new_pwd2")},
			},
		},
	}

	err = clientTest.EditFrontend("created", f, "", version)
	if err != nil {
		t.Error(err.Error())
	} else {
		version++
	}

	v, frontend, err = clientTest.GetFrontend("created", "")
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(frontend, f) {
		fmt.Printf("Edited frontend: %v\n", frontend)
		fmt.Printf("Given frontend: %v\n", f)
		t.Error("Edited frontend not equal to given frontend")
	}

	if v != version {
		t.Errorf("Version %v returned, expected %v", v, version)
	}

	// TestDeleteFrontend
	err = clientTest.DeleteFrontend("created", "", version)
	if err != nil {
		t.Error(err.Error())
	} else {
		version++
	}

	if v, _ := clientTest.GetVersion(""); v != version {
		t.Error("Version not incremented")
	}

	err = clientTest.DeleteFrontend("created", "", 999999)
	if err != nil {
		if confErr, ok := err.(*ConfError); ok {
			if confErr.Code() != ErrVersionMismatch {
				t.Error("Should throw ErrVersionMismatch error")
			}
		} else {
			t.Error("Should throw ErrVersionMismatch error")
		}
	}
	_, _, err = clientTest.GetFrontend("created", "")
	if err == nil {
		t.Error("DeleteFrontend failed, frontend test still exists")
	}

	err = clientTest.DeleteFrontend("doesnotexist", "", version)
	if err == nil {
		t.Error("Should throw error, non existant frontend")
		version++
	}
}
