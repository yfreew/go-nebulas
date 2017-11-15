// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package main

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/nebulasio/go-nebulas/core/state"
	"github.com/nebulasio/go-nebulas/nf/nvm"
	"github.com/nebulasio/go-nebulas/storage"
	"github.com/nebulasio/go-nebulas/util/logging"

	log "github.com/sirupsen/logrus"
)

func main() {
	logging.EnableFuncNameLogger()

	data, _ := ioutil.ReadFile(os.Args[1])

	mem, _ := storage.NewMemoryStorage()
	context, _ := state.NewAccountState(nil, mem)
	owner := context.GetOrCreateUserAccount([]byte("account1"))
	contract, _ := context.CreateContractAccount([]byte("account2"), nil)

	ctx := nvm.NewContext(nil, owner, contract, context)
	engine := nvm.NewV8Engine(ctx)
	err := engine.RunScriptSource(string(data))

	log.Errorf("Err is %s", err)

	time.Sleep(10 * time.Second)
	engine.Dispose()
}
