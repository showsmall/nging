/*
   Nging is a toolbox for webmasters
   Copyright (C) 2018-present Wenhui Shen <swh@admpub.com>

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cloud

import (
	"context"

	"gocloud.dev/blob"
)

var DefaultConfig = &Config{}

type Config struct {
	CloudServerURL string
	PublicBaseURL  string
}

func (c *Config) New(ctx context.Context) (b *blob.Bucket, err error) {
	// Connect to a bucket using a URL, using the "prefix" query parameter to
	// target a subfolder in the bucket.
	// The prefix should end with "/", so that the resulting bucket operates
	// in a subfolder.
	b, err = blob.OpenBucket(ctx, c.CloudServerURL)
	return
}