/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ledgerconfig

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger/testutil"
	"github.com/spf13/viper"
)

func TestIsCouchDBEnabledDefault(t *testing.T) {
	setUpCoreYAMLConfig()
	defaultValue := IsCouchDBEnabled()
	testutil.AssertEquals(t, defaultValue, false) //test default config is false
}

func TestIsCouchDBEnabled(t *testing.T) {
	setUpCoreYAMLConfig()
	defer testutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.stateDatabase", "CouchDB")
	updatedValue := IsCouchDBEnabled()
	testutil.AssertEquals(t, updatedValue, true) //test config returns true
}

func TestGetCouchDBDefinition(t *testing.T) {
	setUpCoreYAMLConfig()
	defer testutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.stateDatabase", "CouchDB")
	couchDBDef := GetCouchDBDefinition()
	testutil.AssertEquals(t, couchDBDef.URL, "127.0.0.1:5984")
	testutil.AssertEquals(t, couchDBDef.Username, "")
	testutil.AssertEquals(t, couchDBDef.Password, "")
}

func TestIsHistoryDBEnabledDefault(t *testing.T) {
	setUpCoreYAMLConfig()
	defaultValue := IsHistoryDBEnabled()
	testutil.AssertEquals(t, defaultValue, false) //test default config is false
}

func TestIsHistoryDBEnabledWhenCouchDBIsDisabled(t *testing.T) {
	setUpCoreYAMLConfig()
	defer testutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.stateDatabase", "goleveldb")
	viper.Set("ledger.state.historyDatabase", true)
	updatedValue := IsHistoryDBEnabled()
	testutil.AssertEquals(t, updatedValue, false) //test config is false
}

func TestIsHistoryDBEnabledWhenOnlyCouchDBEnabled(t *testing.T) {
	setUpCoreYAMLConfig()
	defer testutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.stateDatabase", "CouchDB")
	viper.Set("ledger.state.historyDatabase", false)
	updatedValue := IsHistoryDBEnabled()
	testutil.AssertEquals(t, updatedValue, false) //test config is false
}

func TestIsHistoryDBEnabled(t *testing.T) {
	setUpCoreYAMLConfig()
	defer testutil.ResetConfigToDefaultValues()
	viper.Set("ledger.state.stateDatabase", "CouchDB")
	viper.Set("ledger.state.historyDatabase", true)
	updatedValue := IsHistoryDBEnabled()
	testutil.AssertEquals(t, updatedValue, true) //test config returns true
}

func setUpCoreYAMLConfig() {
	//call a helper method to load the core.yaml
	testutil.SetupCoreYAMLConfig("./../../../peer")
}
