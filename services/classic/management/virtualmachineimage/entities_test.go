// +build go1.7

package virtualmachineimage

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

import (
	"encoding/xml"
	"testing"
)

const xml1 = `
<VMImage>
	<Name>imgName</Name>
	<Label>PackerMade_Ubuntu_Serv14</Label>
	<Category>User</Category>
	<Description>packer made image</Description>
	<OSDiskConfiguration>
		<Name>OSDisk</Name>
		<HostCaching>ReadWrite</HostCaching>
		<OSState>Generalized</OSState>
		<OS>Linux</OS>
		<MediaLink>https://sa.blob.core.usgovcloudapi.net/images/PackerMade_Ubuntu_Serv14_2015-12-12.vhd</MediaLink>
		<LogicalDiskSizeInGB>30</LogicalDiskSizeInGB>
		<IOType>Standard</IOType>
	</OSDiskConfiguration>
	<DataDiskConfigurations/>
	<ServiceName>PkrSrvf3mz03u4mi</ServiceName>
	<DeploymentName>PkrVMf3mz03u4mi</DeploymentName>
	<RoleName>PkrVMf3mz03u4mi</RoleName>
	<Location>Central US</Location>
	<CreatedTime>2015-12-12T08:59:29.1936858Z</CreatedTime>
	<ModifiedTime>2015-12-12T08:59:29.1936858Z</ModifiedTime>
	<ImageFamily>PackerMade</ImageFamily>
	<RecommendedVMSize>Small</RecommendedVMSize>
	<IsPremium>false</IsPremium>
	<VMImageState>VMImageReadyForUse</VMImageState>
	<RoleStateOnCapture>StoppedVM</RoleStateOnCapture>
	<RoleSizeOnCapture>Small</RoleSizeOnCapture>
</VMImage>`
const xml2 = `
<VMImage>
	<Name>imgName</Name>
	<Label>PackerMade_Ubuntu_Serv14</Label>
	<Category>User</Category>
	<Description>packer made image</Description>
	<OSDiskConfiguration>
		<Name>OSDisk</Name>
		<HostCaching>ReadWrite</HostCaching>
		<OSState>Generalized</OSState>
		<OS>Linux</OS>
		<MediaLink>https://sa.blob.core.usgovcloudapi.net/images/PackerMade_Ubuntu_Serv14_2015-12-12.vhd</MediaLink>
		<LogicalDiskSizeInGB>30</LogicalDiskSizeInGB>
		<IOType>Standard</IOType>
	</OSDiskConfiguration>
	<DataDiskConfigurations>
		<DataDiskConfiguration>
			<Name>DataDisk1</Name>
			<HostCaching>ReadWrite</HostCaching>
			<MediaLink>https://sa.blob.core.usgovcloudapi.net/images/PackerMade_Ubuntu_Serv14_2015-12-12-dd1.vhd</MediaLink>
			<LogicalDiskSizeInGB>31</LogicalDiskSizeInGB>
			<IOType>Standard</IOType>
		</DataDiskConfiguration>
		<DataDiskConfiguration>
			<Name>DataDisk2</Name>
			<HostCaching>ReadWrite</HostCaching>
			<MediaLink>https://sa.blob.core.usgovcloudapi.net/images/PackerMade_Ubuntu_Serv14_2015-12-12-dd2.vhd</MediaLink>
			<LogicalDiskSizeInGB>32</LogicalDiskSizeInGB>
			<IOType>Standard</IOType>
		</DataDiskConfiguration>
	</DataDiskConfigurations>
	<ServiceName>PkrSrvf3mz03u4mi</ServiceName>
	<DeploymentName>PkrVMf3mz03u4mi</DeploymentName>
	<RoleName>PkrVMf3mz03u4mi</RoleName>
	<Location>Central US</Location>
	<CreatedTime>2015-12-12T08:59:29.1936858Z</CreatedTime>
	<ModifiedTime>2015-12-12T08:59:29.1936858Z</ModifiedTime>
	<ImageFamily>PackerMade</ImageFamily>
	<RecommendedVMSize>Small</RecommendedVMSize>
	<IsPremium>false</IsPremium>
	<VMImageState>VMImageReadyForUse</VMImageState>
	<RoleStateOnCapture>StoppedVM</RoleStateOnCapture>
	<RoleSizeOnCapture>Small</RoleSizeOnCapture>
</VMImage>`

func Test_NoDataDisksUnmarshal(t *testing.T) {
	var image VMImage
	if err := xml.Unmarshal([]byte(xml1), &image); err != nil {
		t.Fatal(err)
	}

	check := checker{t}
	check.Equal(0, len(image.DataDiskConfigurations))
}

func Test_DataDiskCountUnmarshal(t *testing.T) {
	var image VMImage
	if err := xml.Unmarshal([]byte(xml2), &image); err != nil {
		t.Fatal(err)
	}

	check := checker{t}
	check.Equal(2, len(image.DataDiskConfigurations))
	check.Equal("DataDisk1", image.DataDiskConfigurations[0].Name)
	check.Equal("DataDisk2", image.DataDiskConfigurations[1].Name)
}

type checker struct{ *testing.T }

func (a *checker) Equal(expected, actual interface{}) {
	if expected != actual {
		a.T.Fatalf("Expected %q, but got %q", expected, actual)
	}
}
