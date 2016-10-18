package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonHTMLDocGen(t *testing.T) {
	doc := "Testing 1 2 3"
	expected := "// Testing 1 2 3\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

func TestListsHTMLDocGen(t *testing.T) {
	doc := "<ul><li>Testing 1 2 3</li> <li>FooBar</li></ul>"
	expected := "//    * Testing 1 2 3\n//    * FooBar\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)

	doc = "<ul> <li>Testing 1 2 3</li> <li>FooBar</li> </ul>"
	expected = "//    * Testing 1 2 3\n//    * FooBar\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)

	// Test leading spaces
	doc = " <ul> <li>Testing 1 2 3</li> <li>FooBar</li> </ul>"
	doc = docstring(doc)
	assert.Equal(t, expected, doc)

	// Paragraph check
	doc = "<ul> <li> <p>Testing 1 2 3</p> </li><li> <p>FooBar</p></li></ul>"
	expected = "//    * Testing 1 2 3\n// \n//    * FooBar\n"
	doc = docstring(doc)
	assert.Equal(t, expected, doc)
}

func TestInlineCodeHTMLDocGen(t *testing.T) {
	doc := "<ul> <li><code>Testing</code>: 1 2 3</li> <li>FooBar</li> </ul>"
	expected := "//    * Testing: 1 2 3\n//    * FooBar\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

func TestInlineCodeInParagraphHTMLDocGen(t *testing.T) {
	doc := "<p><code>Testing</code>: 1 2 3</p>"
	expected := "// Testing: 1 2 3\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

func TestEmptyPREInlineCodeHTMLDocGen(t *testing.T) {
	doc := "<pre><code>Testing</code></pre>"
	expected := "//    Testing\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

/*func TestLI(t *testing.T) {
	doc := "<p>String that contains the ARN of the ACM Certificate with one or more tags that you want to remove. This must be of the form:</p> <p> <code>arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012</code> </p> <p>For more information about ARNs, see <a href=\"http://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html\">Amazon Resource Names (ARNs) and AWS Service Namespaces</a>.</p>"
	expected := "//    Testing\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}*/

func TestParagraph(t *testing.T) {
	doc := "<p>Testing 1 2 3</p>"
	expected := "// Testing 1 2 3\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

func TestComplexListParagraphCode(t *testing.T) {
	doc := "<ul> <li><p><code>FOO</code> Bar</p></li><li><p><code>Xyz</code> ABC</p></li></ul>"
	expected := "//    * FOO Bar\n//    * Xyz ABC\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}

func TestRandom(t *testing.T) {
	doc := `
	<p>If specified, the deployment configuration name can be either one of the predefined configurations provided with AWS CodeDeploy or a custom deployment configuration that you create by calling the create deployment configuration operation.</p> <note> <p>CodeDeployDefault.OneAtATime is the default deployment configuration. It is used if a configuration isn't specified for the deployment or the deployment group.</p> </note> <p>The predefined deployment configurations include the following:</p> <ul> <li> <p> <b>CodeDeployDefault.AllAtOnce</b> attempts to deploy an application revision to as many instances as possible at once. The status of the overall deployment will be displayed as <b>Succeeded</b> if the application revision is deployed to one or more of the instances. The status of the overall deployment will be displayed as <b>Failed</b> if the application revision is not deployed to any of the instances. Using an example of nine instances, CodeDeployDefault.AllAtOnce will attempt to deploy to all nine instances at once. The overall deployment will succeed if deployment to even a single instance is successful; it will fail only if deployments to all nine instances fail. </p> </li> <li> <p> <b>CodeDeployDefault.HalfAtATime</b> deploys to up to half of the instances at a time (with fractions rounded down). The overall deployment succeeds if the application revision is deployed to at least half of the instances (with fractions rounded up); otherwise, the deployment fails. In the example of nine instances, it will deploy to up to four instances at a time. The overall deployment succeeds if deployment to five or more instances succeed; otherwise, the deployment fails. The deployment may be successfully deployed to some instances even if the overall deployment fails.</p> </li> <li> <p> <b>CodeDeployDefault.OneAtATime</b> deploys the application revision to only one instance at a time.</p> <p>For deployment groups that contain more than one instance:</p> <ul> <li> <p>The overall deployment succeeds if the application revision is deployed to all of the instances. The exception to this rule is if deployment to the last instance fails, the overall deployment still succeeds. This is because AWS CodeDeploy allows only one instance at a time to be taken offline with the CodeDeployDefault.OneAtATime configuration.</p> </li> <li> <p>The overall deployment fails as soon as the application revision fails to be deployed to any but the last instance. The deployment may be successfully deployed to some instances even if the overall deployment fails.</p> </li> <li> <p>In an example using nine instances, it will deploy to one instance at a time. The overall deployment succeeds if deployment to the first eight instances is successful; the overall deployment fails if deployment to any of the first eight instances fails.</p> </li> </ul> <p>For deployment groups that contain only one instance, the overall deployment is successful only if deployment to the single instance is successful</p> </li> </ul>
	`

	expected := "//    * FOO Bar\n//    * Xyz ABC\n"
	doc = docstring(doc)

	assert.Equal(t, expected, doc)
}
